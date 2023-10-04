package actor

import (
	"context"
	"crm_service/model"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type RepositoryActorInterface interface {

	//CreateActor(ctx *gin.Context, actor *model.Actor) (model.Actor, int, error)
	//GetActorById(ctx *gin.Context, id uint) (model.Actor, int, error)
	//GetAllActor(ctx *gin.Context, page uint, username string) (uint, uint, int, uint, []model.Actor, int, error)
	//UpdateActorById(ctx *gin.Context, id uint, actor *model.Actor) (model.Actor, int, error)
	//DeleteActorById(ctx *gin.Context, id uint) (int, error)
	//ActivateActorById(ctx *gin.Context, id uint) (int, error)
	//DeactivateActorById(ctx *gin.Context, id uint) (int, error)

	LoginActor(ctx *gin.Context, req RequestActor, actorRepository *model.Actor) (int, error)
}

type Actor struct {
	db *gorm.DB
}

func NewActor(dbCrud *gorm.DB) Actor {
	return Actor{
		db: dbCrud,
	}

}

//
//func (repo Actor) CreateActor(actor *model.Actor) (model.Actor, error) {
//	var existingActor model.Actor
//
//	err := repo.db.First(&existingActor, "username = ?", actor.Username).Error
//	if err == nil {
//		// Username already exists, return an error
//		return model.Actor{}, errors.New("username already taken")
//	}
//
//	// Username does not exist, proceed with creating the actor
//	err = repo.db.Create(actor).Error
//	if err != nil {
//		return model.Actor{}, err
//	}
//	registerApproval := model.RegisterApproval{
//		AdminID:      actor.ID,
//		SuperAdminID: 1,
//		Status:       "deactivate",
//	}
//	err = repo.db.Create(&registerApproval).Error
//	if err != nil {
//		return model.Actor{}, err
//	}
//	return *actor, nil
//}
//
//func (repo Actor) GetActorById(id uint) (model.Actor, error) {
//	var actor model.Actor
//	err := repo.db.Omit("password").First(&actor, "id = ?", id).Error
//	if err != nil {
//		return model.Actor{}, errors.New("actor not found")
//	}
//	return actor, nil
//}
//
//func (repo Actor) GetAllActor(page uint, username string) (uint, uint, int, uint, []model.Actor, error) {
//	var actors []model.Actor
//	var count int64
//	var limit uint = 20
//	var offset = limit * (page - 1)
//	result := repo.db.Model(&model.Actor{}).Count(&count)
//	if result.Error != nil {
//		// Handle the error
//		return 0, 0, 0, 0, nil, result.Error
//	}
//	totalPages := uint(math.Ceil(float64(count) / float64(limit)))
//	err := repo.db.Omit("password").Limit(int(limit)).Offset(int(offset)).Where("username LIKE ?", fmt.Sprint("%", username, "%")).Find(&actors).Error
//	if err != nil {
//		return 0, 0, 0, 0, nil, err
//	}
//	return page, limit, int(count), totalPages, actors, nil
//}
//
//func (repo Actor) UpdateActorById(id uint, updateActor *model.Actor) (model.Actor, error) {
//	var findActorById model.Actor
//	var existingActor model.Actor
//
//	if id == 1 {
//		return model.Actor{}, errors.New("actor is super admin and cannot be updated")
//	}
//
//	err := repo.db.First(&findActorById, "id = ?", id).Error
//	if err != nil {
//		return model.Actor{}, errors.New("actor not found")
//	}
//
//	err = repo.db.Where("username = ?", updateActor.Username).Not("username = ?", findActorById.Username).First(&existingActor).Error
//
//	if err == nil {
//		// Username already exists, return an error
//		return model.Actor{}, errors.New("username already taken")
//	}
//
//	err = repo.db.Model(&model.Actor{}).Where("id = ?", id).Updates(updateActor).Error
//	if err != nil {
//		return model.Actor{}, errors.New("failed to update actor")
//	}
//
//	err = repo.db.First(&findActorById, "id = ?", id).Error
//	if err != nil {
//		return model.Actor{}, errors.New("actor not found")
//	}
//
//	return findActorById, nil
//}
//
//func (repo Actor) DeleteActorById(id uint) error {
//	var actor model.Actor
//	if id == 1 {
//		return errors.New("actor is super admin cannot delete")
//	}
//
//	err := repo.db.First(&actor, "id = ?", id).Error
//	if err != nil {
//		return errors.New("actor not found")
//	}
//	err = repo.db.Delete(&actor, "id = ?", id).Error
//	if err != nil {
//		return errors.New("failed deleted")
//	}
//	return nil
//}
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

func (repo Actor) LoginActor(ctx *gin.Context, req RequestActor, actorRepository *model.Actor) (int, error) {
	var cancel context.CancelFunc
	_, cancel = context.WithTimeout(context.Background(), time.Duration(3000)*time.Millisecond)
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
