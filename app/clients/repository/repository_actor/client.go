package repository_actor

import (
	"context"
	"crm_service/app/clients/connection"
	"crm_service/app/config"
	"crm_service/app/model/model_actor"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type ClientRepositoryActor struct {
	client connection.InterfaceConnection
	conf   *config.Config
}

func NewClientActor(conf *config.Config, con connection.InterfaceConnection) InterfaceRepositoryActor {
	return &ClientRepositoryActor{
		client: con,
		conf:   conf,
	}
}

func (repo *ClientRepositoryActor) CreateActor(ctx context.Context, req *model_actor.RequestActor) (int, error) {
	//timeout
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	var args []interface{}
	args = append(args, req.Username, req.Password, time.Now().Format("20060102150405"), req.Username)
	//query
	queryCreateActor := "INSERT INTO actors( username, password,created_at) SELECT ?,?,? WHERE NOT EXISTS (SELECT username FROM actors WHERE username=?)"
	result := repo.client.GetConnectionDB().WithContext(ctx).Exec(queryCreateActor, args...)

	//check
	if result.Error != nil {
		return http.StatusInternalServerError, errors.New("failed exec query create repository-model_actor")
	} else if result.RowsAffected == 0 {
		// Username does not exist, proceed with creating the model
		return http.StatusInternalServerError, errors.New("username already exists")
	}

	//return
	return http.StatusCreated, nil
}

func (repo *ClientRepositoryActor) GetActorByUsername(ctx context.Context, req *model_actor.RequestActor, actorRepository *model_actor.ModelActor) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()
	var args []interface{}
	args = append(args, req.Username)
	//query
	querySelectActor := "SELECT id, username, role_id, verified, active, created_at, updated_at FROM actors WHERE username=?"
	result := repo.client.GetConnectionDB().WithContext(ctx).Raw(querySelectActor, args...).Scan(&actorRepository)
	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query login repository-model_actor")
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("actor not found")
	}

	return http.StatusOK, nil
}

func (repo *ClientRepositoryActor) GetActorById(ctx context.Context, ID uint64) (status int, err error, actorRepository model_actor.ModelActor) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	//query
	queryGetActorById := "SELECT id, username, role_id, verified, active, created_at, updated_at FROM actors WHERE id=?"
	result := repo.client.GetConnectionDB().WithContext(ctx).Raw(queryGetActorById, ID).Scan(&actorRepository)
	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query login repository-model_actor"), actorRepository
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("actor not found"), actorRepository
	}

	return http.StatusOK, nil, actorRepository
}

func (repo *ClientRepositoryActor) GetAllActor(ctx context.Context, page uint64, limit uint64, username string) (status int, err error, actorRepository []model_actor.ModelActor) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	//page
	startID := (page - 1) * limit
	var args []interface{}
	args = append(args, startID, fmt.Sprint("%", username, "%"), limit)

	//query
	queryGetActorById := "SELECT id, username, role_id, verified, active, created_at, updated_at FROM actors WHERE id > ? AND username LIKE ? LIMIT ?"
	result := repo.client.GetConnectionDB().WithContext(ctx).Raw(queryGetActorById, args...).Scan(&actorRepository)

	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query all repository-model_actor"), actorRepository
	}

	return http.StatusOK, nil, actorRepository
}

func (repo *ClientRepositoryActor) GetCountRowsActor(ctx context.Context) (status int, err error, actorRepository model_actor.ModelActor) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	//query
	queryGetActorById := "SELECT count(id) AS total FROM actors"
	result := repo.client.GetConnectionDB().WithContext(ctx).Raw(queryGetActorById).Scan(&actorRepository)
	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query count all repository-model_actor"), actorRepository
	}

	return http.StatusOK, nil, actorRepository
}

func (repo *ClientRepositoryActor) UpdateActorById(ctx context.Context, id uint64, updateActor model_actor.RequestUpdateActor) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	var args []interface{}
	args = append(args, updateActor.Username, updateActor.Verified, updateActor.Active, time.Now().Format("20060102150405"), id)
	//query
	queryUpdateActorById := "UPDATE actors SET username=?,verified=?,active=?,updated_at=? WHERE id=?"
	result := repo.client.GetConnectionDB().WithContext(ctx).Exec(queryUpdateActorById, args...)
	if result.Error != nil {
		//username already exist
		if strings.Contains(result.Error.Error(), "Error 1062 (23000)") {
			return http.StatusBadRequest, errors.New("username already exist")
		}
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query UpdateActorById")
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("data not found")
	}
	return http.StatusAccepted, nil
}

func (repo *ClientRepositoryActor) DeleteActorById(ctx context.Context, id uint64) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	var args []interface{}
	args = append(args, id)

	queryDeleteActorById := "DELETE FROM actors WHERE id =?"
	result := repo.client.GetConnectionDB().WithContext(ctx).Exec(queryDeleteActorById, args...)
	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query DeleteActorById")
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("actor is not found,delete unacceptable")
	}
	return http.StatusOK, nil
}

func (repo *ClientRepositoryActor) ActivateActorById(ctx context.Context, id uint64) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	var args []interface{}
	args = append(args, time.Now().Format("20060102150405"), id)

	queryActivateActorById := "UPDATE actors SET active='true',updated_at=? where id=? AND (active != 'true' OR active IS NULL) "
	result := repo.client.GetConnectionDB().WithContext(ctx).Exec(queryActivateActorById, args...)
	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query ActivateActorById")
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("actor is not found or data has not changed, update unacceptable")
	}
	return http.StatusOK, nil
}

func (repo *ClientRepositoryActor) DeactivateActorById(ctx context.Context, id uint64) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	var args []interface{}
	args = append(args, time.Now().Format("20060102150405"), id)

	queryDeactivateActorById := "UPDATE actors SET active='false',updated_at=? where id=? AND (active != 'false' OR active IS NULL)"
	result := repo.client.GetConnectionDB().WithContext(ctx).Exec(queryDeactivateActorById, args...)
	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query DeactivateActorById")
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("actor is not found or data has not changed, update unacceptable")
	}
	return http.StatusOK, nil
}
