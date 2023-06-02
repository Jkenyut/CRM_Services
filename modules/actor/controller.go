package actor

import (
	"crm_service/dto"
	"crm_service/entity"
	"errors"
)

type ActorControllerInterface interface {
	CreateActor(req ActorBody) (any, error)
	GetActorById(id uint) (FindActor, error)
	GetAllActor() (FindAllActor, error)
	UpdateById(id uint, req UpdateActorBody) (FindActor, error)
	DeleteActorById(id uint) (dto.ResponseMeta, error)
}

type actorControllerStruct struct {
	actorUseCase UseCaseActorInterface
}

func (c actorControllerStruct) CreateActor(req ActorBody) (any, error) {
	actor, err := c.actorUseCase.CreateActor(req)
	if err != nil {
		return SuccessCreate{}, err
	}

	res := SuccessCreate{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success create actor",
			Message:      "Success Register",
			ResponseTime: "",
		},
		Data: ResponseActorBody{
			Username: actor.Username,
		},
	}
	return res, nil
}

func (c actorControllerStruct) GetActorById(id uint) (FindActor, error) {
	var res FindActor
	actor, err := c.actorUseCase.GetActorById(id)
	if err != nil {
		return FindActor{}, err
	}

	if actor.ID == 0 {
		return FindActor{}, errors.New("actor not found")
	}

	res.ResponseMeta = dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Success find actor",
		Message:      "Success Find",
		ResponseTime: "",
	}
	res.Data = actor
	return res, nil
}

func (c actorControllerStruct) GetAllActor() (FindAllActor, error) {
	actorEntities, err := c.actorUseCase.GetAllActor()
	if err != nil {
		return FindAllActor{}, err
	}

	data := make([]entity.Actor, len(actorEntities))
	for i, actorEntity := range actorEntities {
		data[i] = actorEntity
	}

	res := FindAllActor{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success find actor",
			Message:      "Success find all",
			ResponseTime: "",
		},
		Data: data,
	}

	return res, nil
}

func (c actorControllerStruct) UpdateById(id uint, req UpdateActorBody) (FindActor, error) {
	actor, err := c.actorUseCase.UpdateActorById(id, req)
	if actor.ID == 0 {
		return FindActor{}, errors.New("actor not found")
	}

	if err != nil {
		return FindActor{}, err
	}

	res := FindActor{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success update actor",
			Message:      "Success update",
			ResponseTime: "",
		},
		Data: actor,
	}
	return res, nil
}

func (c actorControllerStruct) DeleteActorById(id uint) (dto.ResponseMeta, error) {
	err := c.actorUseCase.DeleteActorById(id)
	res := dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Success delete actor",
		Message:      "Success delete",
		ResponseTime: "",
	}
	return res, err

}
