package actor

import (
	"context"
	"crm_service/dto"
	"crm_service/model"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
)

type ControllerActorInterface interface {
	CreateActor(ctx context.Context, req RequestActor) (dto.DefaultResponse, int, dto.DefaultResponse)
	GetActorById(ctx context.Context, id uint64) (dto.DefaultResponse, int, dto.DefaultResponse)
	GetAllActor(ctx context.Context, page uint64, username string) (dto.DefaultResponse, int, dto.DefaultResponse)
	UpdateActorById(ctx context.Context, id uint64, req RequestUpdateActor) (dto.DefaultResponse, int, dto.DefaultResponse)
	DeleteActorById(ctx context.Context, id uint64) (dto.DefaultResponse, int, dto.DefaultResponse)
	//ActivateActorById(id uint) (dto.ResponseMeta, error)
	//DeactivateActorById(id uint) (dto.ResponseMeta, error)

	LoginActor(ctx context.Context, req RequestActor, agent string) (dto.DefaultResponse, int, dto.DefaultResponse)
}

type actorControllerStruct struct {
	actorRepository RepositoryActorInterface
}

func (c actorControllerStruct) CreateActor(ctx context.Context, req RequestActor) (dto.DefaultResponse, int, dto.DefaultResponse) {
	var actorRepo model.Actor
	var response dto.DefaultResponse
	var errorMessage dto.DefaultResponse

	//hashing password
	hashingPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	reqActor := RequestActor{
		Username: req.Username,
		Password: string(hashingPassword),
	}

	// create actor
	status, err := c.actorRepository.CreateActor(ctx, reqActor)
	if err != nil {
		errorMessage = dto.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	//get data
	status, err = c.actorRepository.GetActorByUsername(ctx, reqActor, &actorRepo)
	if err != nil {
		errorMessage = dto.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	//req approval
	reqApproval := RequestApproval{
		ID: actorRepo.ID,
	}

	//create approval
	status, err = c.actorRepository.CreateApproval(ctx, &reqApproval)
	if err != nil {
		errorMessage = dto.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	response = dto.DefaultSuccessResponseWithMessage("actor created", status, actorRepo)
	return response, status, errorMessage
}

func (c actorControllerStruct) GetActorById(ctx context.Context, id uint64) (dto.DefaultResponse, int, dto.DefaultResponse) {
	var actorRepo model.Actor
	var response dto.DefaultResponse
	var errorMessage dto.DefaultResponse

	//get data by id
	status, err := c.actorRepository.GetActorById(ctx, id, &actorRepo)
	if err != nil {
		errorMessage = dto.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}
	response = dto.DefaultSuccessResponseWithMessage("Get actor", status, actorRepo)
	return response, status, errorMessage
}

func (c actorControllerStruct) GetAllActor(ctx context.Context, page uint64, username string) (dto.DefaultResponse, int, dto.DefaultResponse) {
	var actorRepo []model.Actor
	var actorCountRepo model.Actor
	var response dto.DefaultResponse
	var errorMessage dto.DefaultResponse

	var limit uint64 = 30
	status, err := c.actorRepository.GetCountRowsActor(ctx, &actorCountRepo)
	if err != nil {
		errorMessage = dto.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	status, err = c.actorRepository.GetAllActor(ctx, page, limit, username, &actorRepo)
	if err != nil {
		errorMessage = dto.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	resMessage := FindAllActor{
		Page:       page,
		PerPage:    uint64(len(actorRepo)),
		TotalPages: uint64(math.Ceil(float64(actorCountRepo.Total) / float64(limit))),
		Data:       actorRepo,
	}

	response = dto.DefaultSuccessResponseWithMessage("Get all actor", status, resMessage)

	return response, status, errorMessage
}

func (c actorControllerStruct) UpdateActorById(ctx context.Context, id uint64, req RequestUpdateActor) (dto.DefaultResponse, int, dto.DefaultResponse) {
	var actorRepo model.Actor
	var response dto.DefaultResponse
	var errorMessage dto.DefaultResponse

	//check authorization
	if id == 1 {
		errorMessage = dto.DefaultErrorResponseWithMessage("not authorization update", http.StatusUnauthorized)
		return response, http.StatusUnauthorized, errorMessage
	}

	//update data by id
	status, err := c.actorRepository.UpdateActorById(ctx, id, req)
	if err != nil {
		errorMessage = dto.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	//get data by id
	status, err = c.actorRepository.GetActorById(ctx, id, &actorRepo)
	if err != nil {
		errorMessage = dto.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	response = dto.DefaultSuccessResponseWithMessage("Get actor", status, actorRepo)
	return response, status, errorMessage
}

func (c actorControllerStruct) DeleteActorById(ctx context.Context, id uint64) (dto.DefaultResponse, int, dto.DefaultResponse) {
	var response dto.DefaultResponse
	var errorMessage dto.DefaultResponse
	status, err := c.actorRepository.DeleteActorById(ctx, id)
	if err != nil {
		errorMessage = dto.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	response = dto.DefaultSuccessResponseWithMessage("delete actor", status, "")
	return response, status, errorMessage
}

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

func (c actorControllerStruct) LoginActor(ctx context.Context, req RequestActor, agent string) (dto.DefaultResponse, int, dto.DefaultResponse) {
	var actorRepo model.Actor
	var response dto.DefaultResponse
	var errorMessage dto.DefaultResponse

	status, err := c.actorRepository.LoginActor(ctx, req, &actorRepo)
	if err != nil {
		errorMessage = dto.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	err = bcrypt.CompareHashAndPassword([]byte(actorRepo.Password), []byte(req.Password))
	if err != nil {
		// invalid password
		errorMessage = dto.DefaultErrorResponseWithMessage("invalid username & password", status)
		return response, http.StatusUnauthorized, errorMessage
	}

	//check access
	if actorRepo.Verified != "true" && actorRepo.Active != "true" {
		errorMessage = dto.DefaultErrorResponseWithMessage("account not activated", status)
		return response, http.StatusForbidden, errorMessage
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
		errorMessage = dto.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, http.StatusInternalServerError, errorMessage
	}

	response = dto.DefaultSuccessResponseWithMessage("login success", status, tokenString)
	return response, status, errorMessage
}
