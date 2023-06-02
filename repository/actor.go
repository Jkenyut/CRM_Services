package repository

import (
	"crm_service/entity"
	"errors"
	"gorm.io/gorm"
)

type ActorRepoInterface interface {
	CreateActor(actor *entity.Actor) (*entity.Actor, error)
	GetActorById(id uint) (entity.Actor, error)
	GetAllActor() ([]entity.Actor, error)
	UpdateActorById(id uint, actor *entity.Actor) (entity.Actor, error)
	DeleteActorById(id uint) error
}

type Actor struct {
	db *gorm.DB
}

func NewActor(dbCrud *gorm.DB) Actor {
	return Actor{
		db: dbCrud,
	}

}

func (repo Actor) CreateActor(actor *entity.Actor) (*entity.Actor, error) {
	var existingActor entity.Actor

	err := repo.db.First(&existingActor, "username = ?", actor.Username).Error
	if err == nil {
		// Username already exists, return an error
		return nil, errors.New("username already taken")
	}

	// Username does not exist, proceed with creating the actor
	err = repo.db.Create(actor).Error
	if err != nil {
		return nil, err
	}
	registerApproval := entity.RegisterApproval{
		AdminID:      actor.ID,
		SuperAdminID: 1,
		Status:       "deactivate",
	}
	err = repo.db.Create(&registerApproval).Error
	if err != nil {
		return nil, err
	}
	return actor, nil
}

func (repo Actor) GetActorById(id uint) (entity.Actor, error) {
	var actor entity.Actor
	err := repo.db.First(&actor, "id = ?", id).Error
	if err != nil {
		return entity.Actor{}, errors.New("actor not found")
	}
	return actor, nil
}

func (repo Actor) GetAllActor() ([]entity.Actor, error) {
	var actors []entity.Actor
	err := repo.db.Find(&actors).Error
	if err != nil {
		return nil, err
	}
	return actors, nil
}

func (repo Actor) UpdateActorById(id uint, updateActor *entity.Actor) (entity.Actor, error) {
	var findActor entity.Actor
	if id == 1 {
		return entity.Actor{}, errors.New("actor is super admin and cannot be updated")
	}

	err := repo.db.First(&findActor, "id = ?", id).Error
	if err != nil {
		return entity.Actor{}, errors.New("actor not found")
	}

	err = repo.db.Model(&entity.Actor{}).Where("id = ?", id).Updates(updateActor).Error
	if err != nil {
		return entity.Actor{}, errors.New("failed to update actor")
	}

	err = repo.db.First(&findActor, "id = ?", id).Error
	if err != nil {
		return entity.Actor{}, errors.New("actor not found")
	}

	return findActor, nil
}

func (repo Actor) DeleteActorById(id uint) error {
	var actor entity.Actor
	if id == 1 {
		return errors.New("actor is super admin cannot delete")
	}
	err := repo.db.Delete(&actor, "id = ?", id).Error
	if err != nil {
		return errors.New("actor not found")
	}
	return nil
}
