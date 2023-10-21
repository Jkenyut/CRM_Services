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

	//query
	queryCreateActor := "INSERT INTO actors( username, password) SELECT ?,? WHERE NOT EXISTS (SELECT username from actors where username=?)"
	result := repo.client.GetConnectionDB().WithContext(ctx).Exec(queryCreateActor, req.Username, req.Password, req.Username)

	//check
	if result.Error != nil {
		return http.StatusInternalServerError, errors.New("failed exec query create repository-model_actor")
	} else if result.RowsAffected == 0 {
		// Username does not exist, proceed with creating the repository-model_actor
		return http.StatusInternalServerError, errors.New("username already exists")
	}

	//return
	return http.StatusCreated, nil
}

func (repo *ClientRepositoryActor) createApproval(ctx context.Context, req *model_actor.RequestApproval) (int, error) {
	// timeout
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	//query
	queryCreateApproval := "insert into register_approval(admin_id) values(?)"
	result := repo.client.GetConnectionDB().WithContext(ctx).Exec(queryCreateApproval, req.ID)
	if result.Error != nil {
		// return an if mysql error
		return http.StatusInternalServerError, errors.New("failed exec query create approval")
	} else if result.RowsAffected == 0 {
		return http.StatusInternalServerError, errors.New("failed exec query insert approval")
	}

	return http.StatusOK, nil
}

func (repo *ClientRepositoryActor) getActorByUsername(ctx context.Context, req model_actor.RequestActor, actorRepository *model_actor.ModelActor) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	//query
	querySelectActor := "select id, username, role_id, verified, active, created_at, updated_at from actors where username=?"
	result := repo.client.GetConnectionDB().WithContext(ctx).Raw(querySelectActor, req.Username).Scan(&actorRepository)
	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query login repository-model_actor")
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("repository-model_actor not found")
	}

	return http.StatusOK, nil
}

func (repo *ClientRepositoryActor) GetActorById(ctx context.Context, id uint64, actorRepository *model_actor.ModelActor) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	//query
	queryGetActorById := "select id, username, role_id, verified, active, created_at, updated_at from actors where id=?"
	result := repo.client.GetConnectionDB().WithContext(ctx).Raw(queryGetActorById, id).Scan(&actorRepository)
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

	//query
	queryGetActorById := "select id, username, role_id, verified, active, created_at, updated_at from actors where id > ? AND username like ? limit ?"
	result := repo.client.GetConnectionDB().WithContext(ctx).Raw(queryGetActorById, startID, fmt.Sprint(username, "%"), limit).Scan(&actorRepository)

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
	queryGetActorById := "select count(id) as total from actors"
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

	//query
	queryUpdateActorById := "update actors set username=?,verified=?,verified=? WHERE id=?"
	result := repo.client.GetConnectionDB().WithContext(ctx).Exec(queryUpdateActorById, updateActor.Username, updateActor.Verified, updateActor.Active, id)
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

	queryDeleteActorById := "delete from actors where id =?"
	result := repo.client.GetConnectionDB().WithContext(ctx).Exec(queryDeleteActorById, id)
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

	queryActivateActorById := "update actors set active='true' where id=?"
	result := repo.client.GetConnectionDB().WithContext(ctx).Exec(queryActivateActorById, id)
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

	queryDeactivateActorById := "update actors set verified='true' where id=?"
	result := repo.client.GetConnectionDB().WithContext(ctx).Exec(queryDeactivateActorById, id)
	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query DeactivateActorById")
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("repository-model_actor is not found,update unacceptable")
	}
	return http.StatusOK, nil

}

func (repo *ClientRepositoryActor) LoginActor(ctx context.Context, req model_actor.RequestActor, actorRepository *model_actor.ModelActor) error {
	fmt.Println("wakka")
	fmt.Println("waktu", repo.conf.Database.Timeout)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	fmt.Println("hanya ini")
	queryLoginActor := "select password,verified,role_id,active from actors where username=?"
	err := repo.client.GetConnectionDB().WithContext(ctx).Raw(queryLoginActor, req.Username).Scan(&actorRepository).Error
	fmt.Println("sampe sini")
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// return if not found
			return errors.New("repository-model_actor not found")
		} else {
			// return an if mysql error
			return errors.New("failed exec query login repository-model_actor")
		}
	}
	//ok
	return nil

}
