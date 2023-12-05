package repository_auth

import (
	"context"
	"crm_service/app/clients/connection"
	"crm_service/app/config"
	"crm_service/app/model/model_actor"
	"crm_service/app/model/origin"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
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

func (repo *ClientAuth) InsertSession(ctx context.Context, activity_id string, tokenRefresh string, expiredRefresh time.Time) (status int, error error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	var args []interface{}
	args = append(args, activity_id, tokenRefresh, expiredRefresh, tokenRefresh)

	queryCreateActor := "INSERT INTO sessions(activity_id,jwt_refresh, expired) SELECT ?,?,? WHERE NOT EXISTS (SELECT jwt_refresh FROM session WHERE jwt_refresh=?)"
	result := repo.client.GetConnectionDB().WithContext(ctx).Exec(queryCreateActor, args...)

	//check
	if result.Error != nil {
		return http.StatusInternalServerError, errors.New("failed exec query Insert Session")
	} else if result.RowsAffected == 0 {
		// Username does not exist, proceed with creating the model
		return http.StatusConflict, errors.New("jwt Already Exist")
	}
	return http.StatusOK, nil
}

func (repo *ClientAuth) CheckSession(ctx context.Context, activity_id string) (status int, out origin.JWTModel, error error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	var args []interface{}
	args = append(args, activity_id)

	queryCreateActor := "SELECT jwt_refresh FROM sessions WHERE activity_id = ?"
	result := repo.client.GetConnectionDB().WithContext(ctx).Raw(queryCreateActor, args...).Scan(&out)

	//check
	if result.Error != nil {
		return http.StatusInternalServerError, out, errors.New("failed exec query Check Session")
	} else if result.RowsAffected == 0 {
		// Username does not exist, proceed with creating the model
		return http.StatusNotFound, out, errors.New("Authentication failed, token is access expired.")
	}
	return http.StatusOK, out, nil
}

func (repo *ClientAuth) GenerateJWTAccessCustom(ctx context.Context, req int, agent string, activityId string) (int, string, error) {
	var tokenJWTAccess string
	var err error

	claimsAccess := origin.CustomClaims{
		Data: model_actor.CustomClaimsJWT{
			Role:       strconv.Itoa(req),
			UserAgent:  agent,
			ActivityId: activityId,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "login",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create the token
	tokenAccess := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsAccess)

	// Sign the token with the secret key
	tokenJWTAccess, err = tokenAccess.SignedString([]byte(repo.conf.JWT.JwtAccess))
	if err != nil {
		return http.StatusBadRequest, tokenJWTAccess, errors.New(err.Error())
	}
	return http.StatusOK, tokenJWTAccess, nil

}

func (repo *ClientAuth) GenerateJWTRefreshCustom(ctx context.Context, req int, agent string, activityId string) (int, string, *jwt.NumericDate, error) {
	var tokenJWTRefresh string
	var err error
	claimsRefresh := origin.CustomClaims{
		Data: model_actor.CustomClaimsJWT{
			Role:       strconv.Itoa(req),
			UserAgent:  agent,
			ActivityId: activityId,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "login",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create the token
	tokenRefresh := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)

	// Sign the token with the secret key
	tokenJWTRefresh, err = tokenRefresh.SignedString([]byte(repo.conf.JWT.JwtRefresh))
	if err != nil {
		return http.StatusBadRequest, tokenJWTRefresh, nil, errors.New(err.Error())
	}

	return http.StatusOK, tokenJWTRefresh, claimsRefresh.ExpiresAt, nil
}
