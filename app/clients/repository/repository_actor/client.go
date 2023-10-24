package repository_actor

import (
	"context"
	"crm_service/app/clients/connection"
	"crm_service/app/config"
	"crm_service/app/model/model_actor"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
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
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	var args []interface{}
	args = append(args, req.Username, req.Password, req.Username)
	//query
	queryCreateActor := "INSERT INTO actors( username, password) SELECT ?,? WHERE NOT EXISTS (SELECT username FROM actors WHERE username=?)"
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

func (repo *ClientRepositoryActor) CreateApproval(ctx context.Context, req *model_actor.RequestApproval) (int, error) {
	// timeout
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()
	var args []interface{}
	args = append(args, req.ID)
	//query
	queryCreateApproval := "INSERT INTO register_approval(admin_id) VALUES(?)"
	result := repo.client.GetConnectionDB().WithContext(ctx).Exec(queryCreateApproval, args...)
	if result.Error != nil {
		// return an if mysql error
		return http.StatusInternalServerError, errors.New("failed exec query create approval")
	} else if result.RowsAffected == 0 {
		return http.StatusInternalServerError, errors.New("failed exec query insert approval")
	}

	return http.StatusOK, nil
}

func (repo *ClientRepositoryActor) GetActorByUsername(ctx context.Context, req model_actor.RequestActor, actorRepository *model_actor.ModelActor) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
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
		return http.StatusNotFound, errors.New("repository-model_actor not found")
	}

	return http.StatusOK, nil
}

func (repo *ClientRepositoryActor) GetActorById(ctx context.Context, ID uint64, actorRepository *model_actor.ModelActor) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	//query
	queryGetActorById := "SELECT id, username, role_id, verified, active, created_at, updated_at FROM actors WHERE id=?"
	result := repo.client.GetConnectionDB().WithContext(ctx).Raw(queryGetActorById, ID).Scan(&actorRepository)
	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query login repository-model_actor")
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("repository-model_actor not found")
	}

	return http.StatusOK, nil
}

func (repo *ClientRepositoryActor) GetAllActor(ctx context.Context, page uint64, limit uint64, username string, actorRepository *[]model_actor.ModelActor) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	//page
	startID := (page - 1) * limit
	var args []interface{}
	args = append(args, startID, fmt.Sprint(username, "%"), limit)

	//query
	queryGetActorById := "SELECT id, username, role_id, verified, active, created_at, updated_at FROM actors WHERE id > ? AND username LIKE ? LIMIT ?"
	result := repo.client.GetConnectionDB().WithContext(ctx).Raw(queryGetActorById, args...).Scan(&actorRepository)

	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query all repository-model_actor")
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("all repository-model_actor not found")
	}

	return http.StatusOK, nil
}
func (repo *ClientRepositoryActor) GetCountRowsActor(ctx context.Context, actorRepository *model_actor.ModelActor) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	//query
	queryGetActorById := "SELECT count(id) AS total FROM actors"
	result := repo.client.GetConnectionDB().WithContext(ctx).Raw(queryGetActorById).Scan(&actorRepository)
	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query count all repository-model_actor")
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("count repository-model_actor not found")
	}

	return http.StatusOK, nil
}

func (repo *ClientRepositoryActor) UpdateActorById(ctx context.Context, id uint64, updateActor model_actor.RequestUpdateActor) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	var args []interface{}
	args = append(args, updateActor.Username, updateActor.Verified, updateActor.Active, id)
	//query
	queryUpdateActorById := "UPDATE actors SET username=?,verified=?,activate=? WHERE id=?"
	result := repo.client.GetConnectionDB().WithContext(ctx).Exec(queryUpdateActorById, args...)
	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query UpdateActorById")
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("cannot update because username already exist")
	}

	return http.StatusOK, nil
}

func (repo *ClientRepositoryActor) DeleteActorById(ctx context.Context, id uint64) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
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
		return http.StatusNotFound, errors.New("repository-model_actor is not found,delete unacceptable")
	}
	return http.StatusOK, nil
}

func (repo *ClientRepositoryActor) ActivateActorById(ctx context.Context, id uint64) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	var args []interface{}
	args = append(args, id)

	queryActivateActorById := "UPDATE actors SET active='true' where id=?"
	result := repo.client.GetConnectionDB().WithContext(ctx).Exec(queryActivateActorById, args...)
	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query ActivateActorById")
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("repository-model_actor is not found,update unacceptable")
	}
	return http.StatusOK, nil
}

func (repo *ClientRepositoryActor) DeactivateActorById(ctx context.Context, id uint64) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	var args []interface{}
	args = append(args, id)

	queryDeactivateActorById := "UPDATE actors SET active='false' where id=?"
	result := repo.client.GetConnectionDB().WithContext(ctx).Exec(queryDeactivateActorById, args...)
	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query DeactivateActorById")
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("repository-model_actor is not found,update unacceptable")
	}
	return http.StatusOK, nil

}

func (repo *ClientRepositoryActor) LoginActor(ctx context.Context, req model_actor.RequestActor, actorRepository *model_actor.ModelActor) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	var args []interface{}
	args = append(args, req.Username)

	queryLoginActor := "SELECT password,verified,role_id,active FROM actors WHERE username=?"
	err := repo.client.GetConnectionDB().WithContext(ctx).Raw(queryLoginActor, args...).Scan(&actorRepository).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// return if not found
			return http.StatusNotFound, errors.New("repository-model_actor not found")
		} else {
			// return an if mysql error
			return http.StatusInternalServerError, errors.New("failed exec query login repository-model_actor")
		}
	}
	//ok
	return http.StatusOK, nil

}
