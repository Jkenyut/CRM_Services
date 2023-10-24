package contoller_actor

import (
	"crm_service/app/clients/repository/repository_actor"
	"crm_service/app/model/model_actor"
	"crm_service/app/model/original"
	"crm_service/app/utils/helper"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"strconv"
	"time"
)

type InterfaceControllerActor interface {
	LoginActor(c *gin.Context)
	CreateActor(c *gin.Context)
}

type ControllerActor struct {
	client    repository_actor.InterfaceRepositoryActor
	validator *validator.Validate
}

func NewControllerActor(client repository_actor.InterfaceRepositoryActor, validate *validator.Validate) InterfaceControllerActor {
	return &ControllerActor{
		client:    client,
		validator: validate,
	}
}

func (ctr *ControllerActor) CreateActor(c *gin.Context) {
	// get environment
	//envJWT, _ := c.Get("envJWT")
	//setJWT := envJWT.(map[string]interface{})
	//
	//if setJWT["role"] != "1" {
	//	c.AbortWithStatusJSON(http.StatusUnauthorized, original.DefaultErrorResponseWithMessage("Account Not Authorization", http.StatusUnauthorized))
	//	return
	//}

	// bind to json
	var request model_actor.RequestActor
	err := c.Bind(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, original.DefaultErrorResponseWithMessage("required not valid", http.StatusBadRequest))
		return
	}

	//validate
	err = ctr.validator.Struct(request)
	if err != nil {
		// Validation failed
		c.AbortWithStatusJSON(helper.RequestValidate(err))
		return
	}

	//controllers

	//hashing password
	hashingPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	request.Password = string(hashingPassword)

	var status int
	// create repository-model_actor
	status, err = ctr.client.CreateActor(c, &request)
	if status < 200 || status > 299 {
		c.AbortWithStatusJSON(status, original.DefaultErrorResponseWithMessage(err.Error(), status))
		return
	}

	var responseActor model_actor.ResponseActor
	responseActor.Username = request.Username
	responseActor.CreatedAt = helper.ConvertTimeToWIB(time.Now())
	responseActor.RoleID = 2
	responseActor.Active = "false"

	c.JSON(http.StatusCreated, original.DefaultSuccessResponseWithMessage("actor created", status, responseActor))
}

func (ctr *ControllerActor) GetActorById(c *gin.Context) {
	var actorRepo model_actor.ModelActor
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, original.DefaultErrorResponseWithMessage("must unsigned number", http.StatusBadRequest))
		return
	}
	status, err := ctr.client.GetActorById(c, id, &actorRepo)
	//check status
	if status < 200 || status > 299 {
		c.AbortWithStatusJSON(status, original.DefaultErrorResponseWithMessage(err.Error(), status))
		return
	}

	c.JSON(http.StatusCreated, original.DefaultSuccessResponseWithMessage("actor created", status, actorRepo))
}

//	func (h ControllerActor) GetAllActor(c *gin.Context) {
//		page, err := strconv.ParseUint(c.DefaultQuery("page", "1"), 10, 64)
//		username := c.DefaultQuery("username", "")
//
//		if err != nil {
//			c.AbortWithStatusJSON(http.StatusBadRequest, original.DefaultErrorResponseWithMessage("must unsigned number", http.StatusBadRequest))
//			return
//		}
//
//		res, status, errMessage := h.ctr.GetAllActor(c, page, username)
//		//check status
//		if status < 200 || status > 299 {
//			c.AbortWithStatusJSON(status, errMessage)
//			return
//		}
//		c.JSON(http.StatusOK, res)
//	}
func (ctr *ControllerActor) UpdateActorById(c *gin.Context) {
	//envJWT, _ := c.Get("envJWT")
	//setJWT := envJWT.(map[string]interface{})
	//
	//if setJWT["role"] != "1" {
	//	c.AbortWithStatusJSON(http.StatusUnauthorized, "Not Authorization")
	//	return
	//}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, original.DefaultErrorResponseWithMessage("required not valid", http.StatusBadRequest))
		return
	}

	var request model_actor.RequestUpdateActor
	err = c.Bind(&request)
	//validate
	err = ctr.validator.Struct(request)
	if err != nil {
		// Validation failed
		c.AbortWithStatusJSON(helper.RequestValidate(err))
		return
	}

	status, errMessage := ctr.client.UpdateActorById(c, id, request)
	//check status
	if status < 200 || status > 299 {
		c.AbortWithStatusJSON(status, errMessage)
		return
	}
	var responseActor model_actor.ResponseActor
	responseActor.Username = request.Username
	responseActor.UpdatedAt = helper.ConvertTimeToWIB(time.Now())
	responseActor.RoleID = 2
	responseActor.Verified = request.Verified
	responseActor.Active = request.Active

	c.JSON(http.StatusOK, original.DefaultSuccessResponseWithMessage("Update data actor", status, responseActor))
}

//
//	func (h ControllerActor) DeleteActorById(c *gin.Context) {
//		envJWT, _ := c.Get("envJWT")
//		setJWT := envJWT.(map[string]interface{})
//
//		if setJWT["role"] != "1" {
//			c.AbortWithStatusJSON(http.StatusUnauthorized, "Not Authorization")
//			return
//		}
//		actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
//		if err != nil {
//			c.AbortWithStatusJSON(http.StatusBadRequest, original.DefaultErrorResponseWithMessage("required not valid", http.StatusBadRequest))
//			return
//		}
//
//		res, status, errMessage := h.ctr.DeleteActorById(c, actorId)
//		//check status
//		if status < 200 || status > 299 {
//			c.AbortWithStatusJSON(status, errMessage)
//			return
//		}
//		c.JSON(http.StatusOK, res)
//	}
//
//	func (h ControllerActor) ActivateActorById(c *gin.Context) {
//		envJWT, _ := c.Get("envJWT")
//		setJWT := envJWT.(map[string]interface{})
//
//		if setJWT["role"] != "1" {
//			c.AbortWithStatusJSON(http.StatusUnauthorized, "Not Authorization")
//			return
//		}
//		actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
//		if err != nil {
//			c.JSON(http.StatusBadRequest, original.DefaultErrorResponseWithMessage("required not valid", http.StatusBadRequest))
//			return
//		}
//
//		res, status, errMessage := h.ctr.ActivateActorById(c, actorId)
//		//check status
//		if status < 200 || status > 299 {
//			c.AbortWithStatusJSON(status, errMessage)
//			return
//		}
//		c.JSON(http.StatusOK, res)
//	}
//
//	func (h ControllerActor) DeactivateActorById(c *gin.Context) {
//		envJWT, _ := c.Get("envJWT")
//		setJWT := envJWT.(map[string]interface{})
//
//		if setJWT["role"] != "1" {
//			c.AbortWithStatusJSON(http.StatusUnauthorized, "Not Authorization")
//			return
//		}
//		actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
//		if err != nil {
//			c.JSON(http.StatusBadRequest, original.DefaultErrorResponseWithMessage("required not valid", http.StatusBadRequest))
//			return
//		}
//
//		res, status, errMessage := h.ctr.DeactivateActorById(c, actorId)
//		//check status
//		if status < 200 || status > 299 {
//			c.AbortWithStatusJSON(status, errMessage)
//			return
//		}
//		c.JSON(http.StatusOK, res)
//	}

func (ctr *ControllerActor) LoginActor(c *gin.Context) {
	var actorRepo model_actor.ModelActor
	var response original.DefaultResponse
	var errorMessage original.DefaultResponse
	var status int
	var tokenJWTAccess string

	// get header user-agent
	agent := c.GetHeader("User-Agent")

	//bind to json
	var request model_actor.RequestActor
	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, original.DefaultErrorResponseWithMessage("required not valid", http.StatusBadRequest))
		return
	}

	//logic
	status, err = ctr.client.LoginActor(c, request, &actorRepo)
	if status < 200 || status > 299 {
		errorMessage = original.DefaultErrorResponseWithMessage(err.Error(), status)
		c.AbortWithStatusJSON(status, errorMessage)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(actorRepo.Password), []byte(request.Password))
	if err != nil {
		// invalid password
		errorMessage = original.DefaultErrorResponseWithMessage("invalid username & password", status)
		c.AbortWithStatusJSON(http.StatusUnauthorized, errorMessage)
		return
	}

	//check access
	if actorRepo.Verified != "true" && actorRepo.Active != "true" {
		errorMessage = original.DefaultErrorResponseWithMessage("account not activated", status)
		c.AbortWithStatusJSON(http.StatusForbidden, errorMessage)
		return
	}

	tokenJWTAccess, _, status, err = ctr.GenerateJWTCustom(actorRepo, agent)
	if status < 200 || status > 299 {
		errorMessage = original.DefaultErrorResponseWithMessage(err.Error(), status)
		c.AbortWithStatusJSON(status, errorMessage)
		return
	}

	//response
	response = original.DefaultSuccessResponseWithMessage("login success", status, tokenJWTAccess)

	c.Header("Authorization", "Bearer "+fmt.Sprint(response.Data))
	c.JSON(status, response)
}

func (ctr *ControllerActor) GenerateJWTCustom(req model_actor.ModelActor, agent string) (string, string, int, error) {
	var tokenJWTAccess, tokenJWTRefresh string
	var err error
	claimsRefresh := original.CustomClaims{
		Data: model_actor.CustomClaimsJWT{
			Role:      strconv.Itoa(int(req.RoleID)),
			UserAgent: agent,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "login",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create the token
	tokenRefresh := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)

	// Sign the token with the secret key
	tokenJWTRefresh, err = tokenRefresh.SignedString([]byte(os.Getenv("REFRESH_TOKEN_JWT")))
	if err != nil {
		return tokenJWTAccess, tokenJWTRefresh, http.StatusBadRequest, errors.New(err.Error())
	}

	claimsAccess := original.CustomClaims{
		Data: model_actor.CustomClaimsJWT{
			Role:      strconv.Itoa(int(req.RoleID)),
			UserAgent: agent,
			Refresh:   tokenJWTRefresh,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "login",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create the token
	tokenAccess := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsAccess)

	// Sign the token with the secret key
	tokenJWTAccess, err = tokenAccess.SignedString([]byte(os.Getenv("ACCESS_TOKEN_JWT")))
	if err != nil {
		return tokenJWTAccess, tokenJWTRefresh, http.StatusBadRequest, errors.New(err.Error())
	}
	return tokenJWTAccess, tokenJWTRefresh, http.StatusOK, nil
}

//
//func (h ControllerActor) LogoutActor(c *gin.Context) {
//	//req header and del
//	c.Request.Header.Del("Authorization")
//
//	//response
//	c.JSON(http.StatusOK, original.DefaultSuccessResponseWithMessage("logout success", 200, true))
//}
