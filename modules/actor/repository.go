package actor

import (
	"context"
	"crm_service/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type RepositoryActorInterface interface {
	CreateActor(ctx context.Context, req RequestActor) (int, error)
	CreateApproval(ctx context.Context, req *RequestApproval) (int, error)
	GetActorByUsername(ctx context.Context, req RequestActor, actorRepository *model.Actor) (int, error)
	GetActorById(ctx context.Context, id uint64, actorRepository *model.Actor) (int, error)
	GetAllActor(ctx context.Context, limit uint64, page uint64, username string, actorRepository *[]model.Actor) (int, error)
	GetCountRowsActor(ctx context.Context, actorRepository *model.Actor) (int, error)
	UpdateActorById(ctx context.Context, id uint64, updateActor RequestUpdateActor) (int, error)
	DeleteActorById(ctx context.Context, id uint64) (int, error)
	//ActivateActorById(ctx *context.Context, id uint) (int, error)
	//DeactivateActorById(ctx *context.Context, id uint) (int, error)

	LoginActor(ctx context.Context, req RequestActor, actorRepository *model.Actor) (int, error)
}

type Actor struct {
	db *gorm.DB
}

func NewActor(dbCrud *gorm.DB) Actor {
	return Actor{
		db: dbCrud,
	}

}

func (repo Actor) CreateActor(ctx context.Context, req RequestActor) (int, error) {
	//timeout
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(3000)*time.Millisecond)
	defer cancel()

	//query
	queryCreatenActor := "INSERT INTO actors( username, password) SELECT ?,? WHERE NOT EXISTS (SELECT username from actors where username=?)"
	result := repo.db.WithContext(ctx).Exec(queryCreatenActor, req.Username, req.Password, req.Username)

	//check
	if result.Error != nil {
		return http.StatusInternalServerError, errors.New("failed exec query create actor")
	} else if result.RowsAffected == 0 {
		// Username does not exist, proceed with creating the actor
		return http.StatusInternalServerError, errors.New("username already exists")
	}

	//return
	return http.StatusCreated, nil
}

func (repo Actor) CreateApproval(ctx context.Context, req *RequestApproval) (int, error) {
	// timeout
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(3000)*time.Millisecond)
	defer cancel()

	//query
	queryCreatenApproval := "insert into register_approval(admin_id) values(?)"
	result := repo.db.WithContext(ctx).Exec(queryCreatenApproval, req.ID)
	if result.Error != nil {
		// return an if mysql error
		return http.StatusInternalServerError, errors.New("failed exec query create approval")
	} else if result.RowsAffected == 0 {
		return http.StatusInternalServerError, errors.New("failed exec query insert approval")
	}

	return http.StatusOK, nil
}

func (repo Actor) GetActorByUsername(ctx context.Context, req RequestActor, actorRepository *model.Actor) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(3000)*time.Millisecond)
	defer cancel()

	//query
	querySelectActor := "select id, username, role_id, verified, active, created_at, updated_at from actors where username=?"
	result := repo.db.WithContext(ctx).Raw(querySelectActor, req.Username).Scan(&actorRepository)
	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query login actor")
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("actor not found")
	}

	return http.StatusOK, nil
}

func (repo Actor) GetActorById(ctx context.Context, id uint64, actorRepository *model.Actor) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(3000)*time.Millisecond)
	defer cancel()

	//query
	queryGetActorById := "select id, username, role_id, verified, active, created_at, updated_at from actors where id=?"
	result := repo.db.WithContext(ctx).Raw(queryGetActorById, id).Scan(&actorRepository)
	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query login actor")
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("actor not found")
	}

	return http.StatusOK, nil
}

func (repo Actor) GetAllActor(ctx context.Context, page uint64, limit uint64, username string, actorRepository *[]model.Actor) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(3000)*time.Millisecond)
	defer cancel()

	//page
	startID := (page - 1) * limit

	//query
	queryGetActorById := "select id, username, role_id, verified, active, created_at, updated_at from actors where id > ? AND username like ? limit ?"
	result := repo.db.WithContext(ctx).Raw(queryGetActorById, startID, fmt.Sprint(username, "%"), limit).Scan(&actorRepository)

	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query all actor")
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("all actor not found")
	}

	return http.StatusOK, nil
}
func (repo Actor) GetCountRowsActor(ctx context.Context, actorRepository *model.Actor) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(3000)*time.Millisecond)
	defer cancel()

	//query
	queryGetActorById := "select count(id) as total from actors"
	result := repo.db.WithContext(ctx).Raw(queryGetActorById).Scan(&actorRepository)
	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query login actor")
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("count actor not found")
	}

	return http.StatusOK, nil
}

func (repo Actor) UpdateActorById(ctx context.Context, id uint64, updateActor RequestUpdateActor) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(3000)*time.Millisecond)
	defer cancel()

	//query
	queryUpdateActorById := "update actors set username=?,verified=?,verified=? WHERE id=?"
	result := repo.db.WithContext(ctx).Exec(queryUpdateActorById, updateActor.Username, updateActor.Verified, updateActor.Active, id)
	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query UpdateActorById")
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("cannot update because username already exist")
	}

	return http.StatusOK, nil
}

func (repo Actor) DeleteActorById(ctx context.Context, id uint64) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(3000)*time.Millisecond)
	defer cancel()

	queryDeleteActorById := "delete from actors where id =?"
	result := repo.db.WithContext(ctx).Exec(queryDeleteActorById, id)
	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query DeleteActorById")
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("actor is not found,delete unacceptable")
	}
	return http.StatusOK, nil
}

//
//func (repo Actor) ActivateActorById(id uint) error {
//	var actor model.Actor
//	var register model.RegisterApproval
//
//	err := repo.db.First(&actor, "id = ?", id).Error
//	if err != nil {
//		return errors.New("actor not found")
//	}
//
//	err = repo.db.Model(&register).Where("id = ?", id).Update("status", "activate").Error
//	if err != nil {
//		return errors.New("activate failed")
//	}
//
//	err = repo.db.Model(&actor).Updates(model.Actor{Verified: "true", Active: "true"}).Error
//	if err != nil {
//		return errors.New("activate failed")
//	}
//
//	return nil
//}
//
//func (repo Actor) DeactivateActorById(id uint) error {
//	var actor model.Actor
//	var register model.RegisterApproval
//	if id == 1 {
//		return errors.New("actor is super admin can't deactivate")
//	}
//
//	err := repo.db.First(&actor, "id = ?", id).Error
//	if err != nil {
//		return errors.New("actor not found")
//	}
//
//	err = repo.db.Model(&register).Where("id = ?", id).Update("status", "deactivate").Error
//	if err != nil {
//		return errors.New("deactivate failed")
//	}
//
//	err = repo.db.Model(&actor).Updates(model.Actor{Verified: "false", Active: "false"}).Error
//	if err != nil {
//		return errors.New("deactivate failed")
//	}
//
//	return nil
//}

func (repo Actor) LoginActor(ctx context.Context, req RequestActor, actorRepository *model.Actor) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(3000)*time.Millisecond)
	defer cancel()
	queryLoginActor := "select password,verified,role_id,active from actors where username=?"
	err := repo.db.WithContext(ctx).Raw(queryLoginActor, req.Username).Scan(&actorRepository).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// return if not found
			return http.StatusNotFound, errors.New("actor not found")
		} else {
			// return an if mysql error
			return http.StatusInternalServerError, errors.New("failed exec query login actor")
		}
	}
	//ok
	return http.StatusOK, nil

}
