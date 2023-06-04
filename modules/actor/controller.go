package actor

import (
	"crm_service/dto"
	"crm_service/entity"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"time"
)

type ActorControllerInterface interface {
	CreateActor(req ActorBody) (any, error)
	GetActorById(id uint) (FindActor, error)
	GetAllActor(page uint, usernameStr string) (FindAllActor, error)
	UpdateById(id uint, req UpdateActorBody) (FindActor, error)
	DeleteActorById(id uint) (dto.ResponseMeta, error)
	ActivateActorById(id uint) (dto.ResponseMeta, error)
	DeactivateActorById(id uint) (dto.ResponseMeta, error)
	LoginActor(req ActorBody, agent string) (SuccessLogin, error)
}

type actorControllerStruct struct {
	actorUseCase UseCaseActorInterface
}

func (c actorControllerStruct) CreateActor(req ActorBody) (any, error) {
	start := time.Now()
	actor, err := c.actorUseCase.CreateActor(req)
	if err != nil {
		return SuccessCreate{}, err
	}

	res := SuccessCreate{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success create actor",
			Message:      "Success register",
			ResponseTime: fmt.Sprint(time.Since(start)),
		},
		Data: ActorBody{
			Username: actor.Username,
		},
	}
	return res, nil
}

func (c actorControllerStruct) GetActorById(id uint) (FindActor, error) {
	start := time.Now()
	var res FindActor
	actor, err := c.actorUseCase.GetActorById(id)
	if err != nil {
		return FindActor{}, err
	}

	res.ResponseMeta = dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Success find actor",
		Message:      "Success find",
		ResponseTime: fmt.Sprint(time.Since(start)),
	}
	res.Data = actor
	return res, nil
}

func (c actorControllerStruct) GetAllActor(page uint, usernameStr string) (FindAllActor, error) {
	start := time.Now()
	page, perPage, total, totalPages, actorEntities, err := c.actorUseCase.GetAllActor(page, usernameStr)

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
			ResponseTime: fmt.Sprint(time.Since(start)),
		},
		Page:       page,
		PerPage:    perPage,
		Total:      total,
		TotalPages: totalPages,
		Data:       data,
	}

	return res, nil
}

func (c actorControllerStruct) UpdateById(id uint, req UpdateActorBody) (FindActor, error) {
	start := time.Now()
	actor, err := c.actorUseCase.UpdateActorById(id, req)
	if err != nil {
		return FindActor{}, err
	}

	res := FindActor{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success update actor",
			Message:      "Success update actor",
			ResponseTime: fmt.Sprint(time.Since(start)),
		},
		Data: actor,
	}
	return res, nil
}

func (c actorControllerStruct) DeleteActorById(id uint) (dto.ResponseMeta, error) {
	start := time.Now()
	err := c.actorUseCase.DeleteActorById(id)
	res := dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Success delete actor",
		Message:      "Success delete actor",
		ResponseTime: fmt.Sprint(time.Since(start)),
	}
	return res, err
}

func (c actorControllerStruct) ActivateActorById(id uint) (dto.ResponseMeta, error) {
	start := time.Now()
	err := c.actorUseCase.ActivateActorById(id)
	res := dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Success activate actor",
		Message:      "Success activate actor",
		ResponseTime: fmt.Sprint(time.Since(start)),
	}
	return res, err
}

func (c actorControllerStruct) DeactivateActorById(id uint) (dto.ResponseMeta, error) {
	start := time.Now()
	err := c.actorUseCase.DeactivateActorById(id)
	res := dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Success deactivate actor",
		Message:      "Success deactivate actor",
		ResponseTime: fmt.Sprint(time.Since(start)),
	}
	return res, err
}

func (c actorControllerStruct) LoginActor(req ActorBody, agent string) (SuccessLogin, error) {
	start := time.Now()
	actor, err := c.actorUseCase.LoginActor(req)
	if err != nil {
		return SuccessLogin{}, err
	}

	hour, _ := strconv.Atoi(os.Getenv("HOUR"))
	claims := CustomClaims{Role: uint(actor.RoleID), UserAgent: agent,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Duration(hour) * time.Hour).Unix(),
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_JWT")))
	if err != nil {
		return SuccessLogin{}, errors.New("failed generate to generate token")
	}

	res := SuccessLogin{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success login actor",
			Message:      "Success login actor",
			ResponseTime: fmt.Sprint(time.Since(start)),
		},
		Data: tokenString,
	}
	return res, err
}
