package repository_auth

import (
	"context"
	"crm_service/app/clients/connection"
	"crm_service/app/config"
	"crm_service/app/model/model_actor"
	"crm_service/app/model/origin"
	"errors"
	"fmt"
	"github.com/Jkenyut/libs-numeric-go/libs_models/libs_model_jwt"
	"github.com/Jkenyut/libs-numeric-go/libs_tracing"
	"net/http"
	"time"
)

type ClientAuth struct {
	client connection.InterfaceConnection
	conf   *config.Config
	tr     libs_tracing.InterfaceTracingJaegerOperation
}

func NewClientAuth(conf *config.Config, con connection.InterfaceConnection, tr libs_tracing.InterfaceTracingJaegerOperation) InterfaceAuth {
	return &ClientAuth{
		client: con,
		conf:   conf,
		tr:     tr,
	}
}
func (repo *ClientAuth) LoginActor(ctx context.Context, req model_actor.RequestActor) (actorRepository model_actor.ModelActor, status int, err error) {
	tr := repo.tr
	_, ctx = tr.SetOperationChild("client auth")

	defer tr.FinishChildOperation()

	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	var args []interface{}
	args = append(args, req.Username)
	tr.SetLog("request", req)
	tr.SetLog("args", args)
	queryLoginActor := "SELECT password,verified,role_id,active FROM actors WHERE username=?"

	tr.SetLog("queryLoginActor", queryLoginActor)
	res := repo.client.GetConnectionDB().WithContext(ctx).Raw(queryLoginActor, args...).Scan(&actorRepository)

	if res.Error != nil {
		tr.SetError("err", res.Error.Error())
		// return an if mysql error
		return actorRepository, http.StatusInternalServerError, errors.New("failed exec query login repository-model_actor")
	} else if res.RowsAffected == 0 {
		tr.SetError("err", "Not Found")
		return actorRepository, http.StatusUnauthorized, fmt.Errorf("invalid username & password")
	}
	tr.SetLog("response", actorRepository)

	return actorRepository, http.StatusOK, nil
}

func (repo *ClientAuth) InsertSession(ctx context.Context, activityId string, agent string, claimRefresh *libs_model_jwt.CustomClaims) (status int, error error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	var args []interface{}
	args = append(args, activityId, agent, claimRefresh.IssuedAt.Time, time.Now().Add(time.Duration(repo.conf.JWT.ExpiredRefresh)*time.Hour), activityId)

	queryCreateActor := "INSERT INTO sessions(activity_id,agent,issued_at, expired_at) SELECT ?,?,?,? WHERE NOT EXISTS (SELECT activity_id FROM sessions WHERE activity_id=?)"
	result := repo.client.GetConnectionDB().WithContext(ctx).Exec(queryCreateActor, args...)

	//check
	if result.Error != nil {
		return http.StatusInternalServerError, errors.New("failed exec query Insert Session")
	} else if result.RowsAffected == 0 {
		// Username does not exist, proceed with creating the model
		return http.StatusConflict, errors.New("refresh Already Exist")
	}
	return http.StatusOK, nil
}

func (repo *ClientAuth) CheckSession(ctx context.Context, activityId string) (status int, out origin.JWTModel, error error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	var args []interface{}
	args = append(args, activityId, time.Now())

	queryCreateActor := "SELECT agent,issued_at FROM sessions WHERE activity_id = ? AND expired_at > ?"
	result := repo.client.GetConnectionDB().WithContext(ctx).Raw(queryCreateActor, args...).Scan(&out)

	//check
	if result.Error != nil {
		return http.StatusInternalServerError, out, errors.New("failed exec query Check Session")
	} else if result.RowsAffected == 0 {
		// Username does not exist, proceed with creating the model
		return http.StatusNotFound, out, errors.New(" authentication failed,refresh token is access expired.")
	}
	return http.StatusOK, out, nil
}

func (repo *ClientAuth) DeleteSession(ctx context.Context, activityId string) (status int, error error) {
	tr := repo.tr
	span, ctx := tr.SetOperationChild("client Delete")
	repo.tr.SetLog("HH", "ooo")
	defer span.Finish()

	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	var args []interface{}
	args = append(args, activityId)

	queryDeleteSession := "DELETE FROM sessions WHERE activity_id = ?"
	result := repo.client.GetConnectionDB().WithContext(ctx).Exec(queryDeleteSession, args...)
	//check
	if result.Error != nil {
		return http.StatusInternalServerError, errors.New("failed exec query Delete Session")
	} else if result.RowsAffected == 0 {
		// Username does not exist, proceed with creating the model
		return http.StatusNotFound, errors.New(" authentication failed,refresh token is not found.")
	}

	queryDeleteAllSession := "DELETE FROM sessions WHERE expired_at < ?"
	result = repo.client.GetConnectionDB().WithContext(ctx).Exec(queryDeleteAllSession, time.Now())
	//check
	if result.Error != nil {
		return http.StatusInternalServerError, errors.New("failed exec query Delete ALL Session")
	}

	return http.StatusAccepted, nil
}
