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
	"net/http"
	"time"
)

type ClientAuth struct {
	client connection.InterfaceConnection
	conf   *config.Config
}

func NewClientAuth(conf *config.Config, con connection.InterfaceConnection) InterfaceAuth {
	return &ClientAuth{
		client: con,
		conf:   conf,
	}
}
func (repo *ClientAuth) LoginActor(ctx context.Context, req model_actor.RequestActor, actorRepository *model_actor.ModelActor) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	var args []interface{}
	args = append(args, req.Username)

	queryLoginActor := "SELECT password,verified,role_id,active FROM actors WHERE username=?"
	res := repo.client.GetConnectionDB().WithContext(ctx).Raw(queryLoginActor, args...).Scan(&actorRepository)

	if res.Error != nil {
		// return an if mysql error
		return http.StatusInternalServerError, errors.New("failed exec query login repository-model_actor")
	} else if res.RowsAffected == 0 {
		return http.StatusUnauthorized, fmt.Errorf("invalid username & password")
	}

	return http.StatusOK, nil
}

func (repo *ClientAuth) InsertSession(ctx context.Context, activityId string, agent string, claimRefresh libs_model_jwt.CustomClaims) (status int, error error) {
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
	return http.StatusAccepted, nil
}
