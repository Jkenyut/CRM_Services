package actor

import (
	"context"
	"crm_service/entity"
	"crm_service/middleware"
	"crm_service/model"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
)

type ControllerActorInterface interface {
	CreateActor(ctx context.Context, req RequestActor) (entity.DefaultResponse, int, entity.DefaultResponse)
	GetActorById(ctx context.Context, id uint64) (entity.DefaultResponse, int, entity.DefaultResponse)
	GetAllActor(ctx context.Context, page uint64, username string) (entity.DefaultResponse, int, entity.DefaultResponse)
	UpdateActorById(ctx context.Context, id uint64, req RequestUpdateActor) (entity.DefaultResponse, int, entity.DefaultResponse)
	DeleteActorById(ctx context.Context, id uint64) (entity.DefaultResponse, int, entity.DefaultResponse)
	ActivateActorById(ctx context.Context, id uint64) (entity.DefaultResponse, int, entity.DefaultResponse)
	DeactivateActorById(ctx context.Context, id uint64) (entity.DefaultResponse, int, entity.DefaultResponse)

	LoginActor(ctx context.Context, req RequestActor, agent string) (entity.DefaultResponse, int, entity.DefaultResponse)
}

type actorControllerStruct struct {
	actorRepository RepositoryActorInterface
}

func (c actorControllerStruct) CreateActor(ctx context.Context, req RequestActor) (entity.DefaultResponse, int, entity.DefaultResponse) {
	var actorRepo model.Actor
	var response entity.DefaultResponse
	var errorMessage entity.DefaultResponse

	//hashing password
	hashingPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	reqActor := RequestActor{
		Username: req.Username,
		Password: string(hashingPassword),
	}

	// create actor
	status, err := c.actorRepository.CreateActor(ctx, reqActor)
	if err != nil {
		errorMessage = entity.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	//get data
	status, err = c.actorRepository.GetActorByUsername(ctx, reqActor, &actorRepo)
	if err != nil {
		errorMessage = entity.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	//req approval
	reqApproval := RequestApproval{
		ID: actorRepo.ID,
	}

	//create approval
	status, err = c.actorRepository.CreateApproval(ctx, &reqApproval)
	if err != nil {
		errorMessage = entity.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	response = entity.DefaultSuccessResponseWithMessage("actor created", status, actorRepo)
	return response, status, errorMessage
}

func (c actorControllerStruct) GetActorById(ctx context.Context, id uint64) (entity.DefaultResponse, int, entity.DefaultResponse) {
	var actorRepo model.Actor
	var response entity.DefaultResponse
	var errorMessage entity.DefaultResponse

	//get data by id
	status, err := c.actorRepository.GetActorById(ctx, id, &actorRepo)
	if err != nil {
		errorMessage = entity.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}
	response = entity.DefaultSuccessResponseWithMessage("Get actor", status, actorRepo)
	return response, status, errorMessage
}

func (c actorControllerStruct) GetAllActor(ctx context.Context, page uint64, username string) (entity.DefaultResponse, int, entity.DefaultResponse) {
	var actorRepo []model.Actor
	var actorCountRepo model.Actor
	var response entity.DefaultResponse
	var errorMessage entity.DefaultResponse

	var limit uint64 = 30
	status, err := c.actorRepository.GetCountRowsActor(ctx, &actorCountRepo)
	if err != nil {
		errorMessage = entity.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	status, err = c.actorRepository.GetAllActor(ctx, page, limit, username, &actorRepo)
	if err != nil {
		errorMessage = entity.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	resMessage := FindAllActor{
		Page:       page,
		PerPage:    uint64(len(actorRepo)),
		TotalPages: uint64(math.Ceil(float64(actorCountRepo.Total) / float64(limit))),
		Data:       actorRepo,
	}

	response = entity.DefaultSuccessResponseWithMessage("Get all actor", status, resMessage)

	return response, status, errorMessage
}

func (c actorControllerStruct) UpdateActorById(ctx context.Context, id uint64, req RequestUpdateActor) (entity.DefaultResponse, int, entity.DefaultResponse) {
	var actorRepo model.Actor
	var response entity.DefaultResponse
	var errorMessage entity.DefaultResponse

	//check authorization
	if id == 1 {
		errorMessage = entity.DefaultErrorResponseWithMessage("not authorization update", http.StatusUnauthorized)
		return response, http.StatusUnauthorized, errorMessage
	}

	//update data by id
	status, err := c.actorRepository.UpdateActorById(ctx, id, req)
	if err != nil {
		errorMessage = entity.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	//repo
	status, err = c.actorRepository.GetActorById(ctx, id, &actorRepo)
	if err != nil {
		errorMessage = entity.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	response = entity.DefaultSuccessResponseWithMessage("Get actor", status, actorRepo)
	return response, status, errorMessage
}

func (c actorControllerStruct) DeleteActorById(ctx context.Context, id uint64) (entity.DefaultResponse, int, entity.DefaultResponse) {
	var response entity.DefaultResponse
	var errorMessage entity.DefaultResponse
	//check authorization
	if id == 1 {
		errorMessage = entity.DefaultErrorResponseWithMessage("not authorization delete", http.StatusUnauthorized)
		return response, http.StatusUnauthorized, errorMessage
	}

	//repo
	status, err := c.actorRepository.DeleteActorById(ctx, id)
	if err != nil {
		errorMessage = entity.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}
	response = entity.DefaultSuccessResponseWithMessage("delete actor", status, "true")
	return response, status, errorMessage
}

func (c actorControllerStruct) ActivateActorById(ctx context.Context, id uint64) (entity.DefaultResponse, int, entity.DefaultResponse) {
	var response entity.DefaultResponse
	var errorMessage entity.DefaultResponse

	//repo
	status, err := c.actorRepository.ActivateActorById(ctx, id)
	if err != nil {
		errorMessage = entity.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	response = entity.DefaultSuccessResponseWithMessage("delete actor", status, "true")
	return response, status, errorMessage
}

func (c actorControllerStruct) DeactivateActorById(ctx context.Context, id uint64) (entity.DefaultResponse, int, entity.DefaultResponse) {
	var response entity.DefaultResponse
	var errorMessage entity.DefaultResponse
	//check authorization
	if id == 1 {
		errorMessage = entity.DefaultErrorResponseWithMessage("not authorization deactivate", http.StatusUnauthorized)
		return response, http.StatusUnauthorized, errorMessage
	}

	//repo
	status, err := c.actorRepository.DeactivateActorById(ctx, id)
	if err != nil {
		errorMessage = entity.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}
	response = entity.DefaultSuccessResponseWithMessage("delete actor", status, "true")
	return response, status, errorMessage
}

func (c actorControllerStruct) LoginActor(ctx context.Context, req RequestActor, agent string) (entity.DefaultResponse, int, entity.DefaultResponse) {
	var actorRepo model.Actor
	var response entity.DefaultResponse
	var errorMessage entity.DefaultResponse

	status, err := c.actorRepository.LoginActor(ctx, req, &actorRepo)
	if err != nil {
		errorMessage = entity.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	err = bcrypt.CompareHashAndPassword([]byte(actorRepo.Password), []byte(req.Password))
	if err != nil {
		// invalid password
		errorMessage = entity.DefaultErrorResponseWithMessage("invalid username & password", status)
		return response, http.StatusUnauthorized, errorMessage
	}

	//check access
	if actorRepo.Verified != "true" && actorRepo.Active != "true" {
		errorMessage = entity.DefaultErrorResponseWithMessage("account not activated", status)
		return response, http.StatusForbidden, errorMessage
	}

	claims := middleware.CustomClaims{
		Data: customClaimsJWT{
			Role:      strconv.Itoa(int(actorRepo.RoleID)),
			UserAgent: agent,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "login",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_JWT")))
	if err != nil {
		errorMessage = entity.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, http.StatusInternalServerError, errorMessage
	}

	response = entity.DefaultSuccessResponseWithMessage("login success", status, tokenString)
	return response, status, errorMessage
}
