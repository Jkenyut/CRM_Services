package actor

import (
	"crm_service/entity"
	"crm_service/repository"
	"golang.org/x/crypto/bcrypt"
)

type UseCaseActorInterface interface {
	CreateActor(actor ActorBody) (entity.Actor, error)
	GetActorById(id uint) (entity.Actor, error)
	GetAllActor(page uint) (uint, uint, int, uint, []entity.Actor, error)
	UpdateActorById(id uint, actor UpdateActorBody) (entity.Actor, error)
	DeleteActorById(id uint) error
}

type actorUseCaseStruct struct {
	actorRepository repository.ActorRepoInterface
}

func (uc actorUseCaseStruct) CreateActor(actor ActorBody) (entity.Actor, error) {

	hashingPassword, _ := bcrypt.GenerateFromPassword([]byte(actor.Password), 12)
	NewActor := entity.Actor{
		Username: actor.Username,
		Password: string(hashingPassword),
	}

	createActor, err := uc.actorRepository.CreateActor(&NewActor)
	if err != nil {
		return NewActor, err
	}
	return createActor, nil
}

func (uc actorUseCaseStruct) GetActorById(id uint) (entity.Actor, error) {
	var actor entity.Actor
	actor, err := uc.actorRepository.GetActorById(id)
	return actor, err
}

func (uc actorUseCaseStruct) GetAllActor(page uint) (uint, uint, int, uint, []entity.Actor, error) {
	var actor []entity.Actor
	page, perPage, total, totalPages, actor, err := uc.actorRepository.GetAllActor(page)
	return page, perPage, total, totalPages, actor, err
}

func (uc actorUseCaseStruct) UpdateActorById(id uint, actor UpdateActorBody) (entity.Actor, error) {
	hashingPassword, _ := bcrypt.GenerateFromPassword([]byte(actor.Password), 12)
	newActor := entity.Actor{
		Username: actor.Username,
		Password: string(hashingPassword),
		Verified: actor.Verified,
		Active:   actor.Active,
	}

	updatedActor, err := uc.actorRepository.UpdateActorById(id, &newActor)
	if err != nil {
		return newActor, err
	}

	return updatedActor, nil
}

func (uc actorUseCaseStruct) DeleteActorById(id uint) error {
	err := uc.actorRepository.DeleteActorById(id)
	return err

}
