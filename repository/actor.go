package repository

import (
	"crm_service/entity"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"math"
)

type ActorRepoInterface interface {
	CreateActor(actor *entity.Actor) (entity.Actor, error)
	GetActorById(id uint) (entity.Actor, error)
	GetAllActor(page uint, username string) (uint, uint, int, uint, []entity.Actor, error)
	UpdateActorById(id uint, actor *entity.Actor) (entity.Actor, error)
	DeleteActorById(id uint) error
	ActivateActorById(id uint) error
	DeactivateActorById(id uint) error
	LoginActor(actor *entity.Actor) (entity.Actor, error)
}

type Actor struct {
	db *gorm.DB
}

func NewActor(dbCrud *gorm.DB) Actor {
	return Actor{
		db: dbCrud,
	}

}

func (repo Actor) CreateActor(actor *entity.Actor) (entity.Actor, error) {
	var existingActor entity.Actor

	err := repo.db.First(&existingActor, "username = ?", actor.Username).Error
	if err == nil {
		// Username already exists, return an error
		return entity.Actor{}, errors.New("username already taken")
	}

	// Username does not exist, proceed with creating the actor
	err = repo.db.Create(actor).Error
	if err != nil {
		return entity.Actor{}, err
	}
	registerApproval := entity.RegisterApproval{
		AdminID:      actor.ID,
		SuperAdminID: 1,
		Status:       "deactivate",
	}
	err = repo.db.Create(&registerApproval).Error
	if err != nil {
		return entity.Actor{}, err
	}
	return *actor, nil
}

func (repo Actor) GetActorById(id uint) (entity.Actor, error) {
	var actor entity.Actor
	err := repo.db.Omit("password").First(&actor, "id = ?", id).Error
	if err != nil {
		return entity.Actor{}, errors.New("actor not found")
	}
	return actor, nil
}

func (repo Actor) GetAllActor(page uint, username string) (uint, uint, int, uint, []entity.Actor, error) {
	var actors []entity.Actor
	var count int64
	var limit uint = 20
	var offset = limit * (page - 1)
	result := repo.db.Model(&entity.Actor{}).Count(&count)
	if result.Error != nil {
		// Handle the error
		return 0, 0, 0, 0, nil, result.Error
	}
	totalPages := uint(math.Ceil(float64(count) / float64(limit)))
	err := repo.db.Omit("password").Limit(int(limit)).Offset(int(offset)).Where("username LIKE ?", fmt.Sprint("%", username, "%")).Find(&actors).Error
	if err != nil {
		return 0, 0, 0, 0, nil, err
	}
	return page, limit, int(count), totalPages, actors, nil
}

func (repo Actor) UpdateActorById(id uint, updateActor *entity.Actor) (entity.Actor, error) {
	var findActorById entity.Actor
	var existingActor entity.Actor

	if id == 1 {
		return entity.Actor{}, errors.New("actor is super admin and cannot be updated")
	}

	err := repo.db.First(&findActorById, "id = ?", id).Error
	if err != nil {
		return entity.Actor{}, errors.New("actor not found")
	}

	err = repo.db.Where("username = ?", updateActor.Username).Not("username = ?", findActorById.Username).First(&existingActor).Error

	if err == nil {
		// Username already exists, return an error
		return entity.Actor{}, errors.New("username already taken")
	}

	err = repo.db.Model(&entity.Actor{}).Where("id = ?", id).Updates(updateActor).Error
	if err != nil {
		return entity.Actor{}, errors.New("failed to update actor")
	}

	err = repo.db.First(&findActorById, "id = ?", id).Error
	if err != nil {
		return entity.Actor{}, errors.New("actor not found")
	}

	return findActorById, nil
}

func (repo Actor) DeleteActorById(id uint) error {
	var actor entity.Actor
	if id == 1 {
		return errors.New("actor is super admin cannot delete")
	}

	err := repo.db.First(&actor, "id = ?", id).Error
	if err != nil {
		return errors.New("actor not found")
	}
	err = repo.db.Delete(&actor, "id = ?", id).Error
	if err != nil {
		return errors.New("failed deleted")
	}
	return nil
}

func (repo Actor) ActivateActorById(id uint) error {
	var actor entity.Actor
	var register entity.RegisterApproval

	err := repo.db.First(&actor, "id = ?", id).Error
	if err != nil {
		return errors.New("actor not found")
	}

	err = repo.db.Model(&register).Where("id = ?", id).Update("status", "activate").Error
	if err != nil {
		return errors.New("activate failed")
	}

	err = repo.db.Model(&actor).Updates(entity.Actor{Verified: "true", Active: "true"}).Error
	if err != nil {
		return errors.New("activate failed")
	}

	return nil
}

func (repo Actor) DeactivateActorById(id uint) error {
	var actor entity.Actor
	var register entity.RegisterApproval
	if id == 1 {
		return errors.New("actor is super admin can't deactivate")
	}

	err := repo.db.First(&actor, "id = ?", id).Error
	if err != nil {
		return errors.New("actor not found")
	}

	err = repo.db.Model(&register).Where("id = ?", id).Update("status", "deactivate").Error
	if err != nil {
		return errors.New("deactivate failed")
	}

	err = repo.db.Model(&actor).Updates(entity.Actor{Verified: "false", Active: "false"}).Error
	if err != nil {
		return errors.New("deactivate failed")
	}

	return nil
}

func (repo Actor) LoginActor(actor *entity.Actor) (entity.Actor, error) {
	var existingActor entity.Actor

	err := repo.db.First(&existingActor, "username = ?", actor.Username).Error

	if err != nil {

		// Username not found, return an error
		return entity.Actor{}, errors.New("actor not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingActor.Password), []byte(actor.Password))
	if err != nil {
		// invalid
		return entity.Actor{}, errors.New("invalid username & password")
	}

	if existingActor.Verified != "true" && existingActor.Active != "true" {
		return entity.Actor{}, errors.New("username not activate")
	}
	return existingActor, err

}
