package actor

import (
	"crm_service/dto"
	"crm_service/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"strconv"
	"time"
)

type ControllerActorInterface interface {
	//CreateActor(req ActorRequest) (any, error)
	//GetActorById(id uint) (FindActor, error)
	//GetAllActor(page uint, usernameStr string) (FindAllActor, error)
	//UpdateById(id uint, req UpdateActorRequest) (FindActor, error)
	//DeleteActorById(id uint) (dto.ResponseMeta, error)
	//ActivateActorById(id uint) (dto.ResponseMeta, error)
	//DeactivateActorById(id uint) (dto.ResponseMeta, error)

	LoginActor(ctx *gin.Context, req RequestActor, agent string) (dto.DefaultResponse, int, dto.DefaultResponse)
}

type actorControllerStruct struct {
	actorRepository RepositoryActorInterface
}

//func (c actorControllerStruct) CreateActor(req ActorRequest) (any, error) {
//	start := time.Now()
//	actor, err := c.actorUseCase.CreateActor(req)
//	if err != nil {
//		return SuccessCreate{}, err
//	}
//
//	res := SuccessCreate{
//		ResponseMeta: dto.ResponseMeta{
//			Success:      true,
//			Message:      "Success register",
//			ResponseTime: fmt.Sprint(time.Since(start)),
//		},
//		Data: ActorRequest{
//			Username: actor.Username,
//		},
//	}
//	return res, nil
//}
//
//func (c actorControllerStruct) GetActorById(id uint) (FindActor, error) {
//	start := time.Now()
//	var res FindActor
//	actor, err := c.actorUseCase.GetActorById(id)
//	if err != nil {
//		return FindActor{}, err
//	}
//
//	res.ResponseMeta = dto.ResponseMeta{
//		Success:      true,
//		Message:      "Success find",
//		ResponseTime: fmt.Sprint(time.Since(start)),
//	}
//	res.Data = actor
//	return res, nil
//}
//
//func (c actorControllerStruct) GetAllActor(page uint, usernameStr string) (FindAllActor, error) {
//	start := time.Now()
//	page, perPage, total, totalPages, actorEntities, err := c.actorUseCase.GetAllActor(page, usernameStr)
//
//	if err != nil {
//		return FindAllActor{}, err
//	}
//
//	data := make([]model.Actor, len(actorEntities))
//	for i, actorEntity := range actorEntities {
//		data[i] = actorEntity
//	}
//
//	res := FindAllActor{
//		ResponseMeta: dto.ResponseMeta{
//			Success:      true,
//			Message:      "Success find all",
//			ResponseTime: fmt.Sprint(time.Since(start)),
//		},
//		Page:       page,
//		PerPage:    perPage,
//		Total:      total,
//		TotalPages: totalPages,
//		Data:       data,
//	}
//
//	return res, nil
//}
//
//func (c actorControllerStruct) UpdateById(id uint, req UpdateActorRequest) (FindActor, error) {
//	start := time.Now()
//	actor, err := c.actorUseCase.UpdateActorById(id, req)
//	if err != nil {
//		return FindActor{}, err
//	}
//
//	res := FindActor{
//		ResponseMeta: dto.ResponseMeta{
//			Success:      true,
//			Message:      "Success update actor",
//			ResponseTime: fmt.Sprint(time.Since(start)),
//		},
//		Data: actor,
//	}
//	return res, nil
//}
//
//func (c actorControllerStruct) DeleteActorById(id uint) (dto.ResponseMeta, error) {
//	start := time.Now()
//	err := c.actorUseCase.DeleteActorById(id)
//	res := dto.ResponseMeta{
//		Success:      true,
//		Message:      "Success delete actor",
//		ResponseTime: fmt.Sprint(time.Since(start)),
//	}
//	return res, err
//}
//
//func (c actorControllerStruct) ActivateActorById(id uint) (dto.ResponseMeta, error) {
//	start := time.Now()
//	err := c.actorUseCase.ActivateActorById(id)
//	res := dto.ResponseMeta{
//		Success:      true,
//		Message:      "Success activate actor",
//		ResponseTime: fmt.Sprint(time.Since(start)),
//	}
//	return res, err
//}
//
//func (c actorControllerStruct) DeactivateActorById(id uint) (dto.ResponseMeta, error) {
//	start := time.Now()
//	err := c.actorUseCase.DeactivateActorById(id)
//	res := dto.ResponseMeta{
//		Success:      true,
//		Message:      "Success deactivate actor",
//		ResponseTime: fmt.Sprint(time.Since(start)),
//	}
//	return res, err
//}

func (c actorControllerStruct) LoginActor(ctx *gin.Context, req RequestActor, agent string) (dto.DefaultResponse, int, dto.DefaultResponse) {
	start := time.Now()
	var actorRepo model.Actor
	status, err := c.actorRepository.LoginActor(ctx, req, &actorRepo)
	if err != nil {
		errorMessage := dto.DefaultErrorResponseWithMessage(err.Error(), fmt.Sprint(time.Since(start)), strconv.Itoa(status))
		return dto.DefaultResponse{}, status, errorMessage
	}

	err = bcrypt.CompareHashAndPassword([]byte(actorRepo.Password), []byte(req.Password))
	if err != nil {
		// invalid
		errorMessage := dto.DefaultErrorResponseWithMessage("invalid username & password", fmt.Sprint(time.Since(start)), strconv.Itoa(status))
		return dto.DefaultResponse{}, http.StatusUnauthorized, errorMessage
	}

	if actorRepo.Verified != "true" && actorRepo.Active != "true" {
		errorMessage := dto.DefaultErrorResponseWithMessage("account not activated", fmt.Sprint(time.Since(start)), strconv.Itoa(status))
		return dto.DefaultResponse{}, http.StatusForbidden, errorMessage
	}

	hour, _ := strconv.Atoi(os.Getenv("HOUR"))
	claims := CustomClaims{Role: uint(actorRepo.RoleID), UserAgent: agent,
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
		errorMessage := dto.DefaultErrorResponseWithMessage(err.Error(), fmt.Sprint(time.Since(start)), strconv.Itoa(status))
		return dto.DefaultResponse{}, http.StatusInternalServerError, errorMessage
	}

	res := dto.DefaultSuccessResponseWithMessage("login success", fmt.Sprint(time.Since(start)), strconv.Itoa(status), tokenString)
	return res, status, dto.DefaultResponse{}
}
