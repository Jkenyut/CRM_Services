package controller_auth_actor

import (
	"crm_service/app/clients/repository/repository_auth"
	"crm_service/app/model/model_actor"
	"crm_service/app/model/origin"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type ControllerAuth struct {
	client    repository_auth.InterfaceAuth
	validator *validator.Validate
}

func NewControllerAuth(client repository_auth.InterfaceAuth, validate *validator.Validate) InterfaceControllerAuth {
	return &ControllerAuth{
		client:    client,
		validator: validate,
	}
}

func (ctr *ControllerAuth) LoginActor(c *gin.Context) {
	var actorRepo model_actor.ModelActor
	var response origin.DefaultResponse
	var errorMessage origin.DefaultResponse
	var status int
	var tokenJWTAccess, tokenJWTRefresh string
	var ExpiredRefresh *jwt.NumericDate

	// get header user-agent
	agent := c.GetHeader("User-Agent")

	//bind to json
	var request model_actor.RequestActor
	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, origin.DefaultErrorResponseWithMessage("required not valid", http.StatusBadRequest))
		return
	}

	//logic
	status, err = ctr.client.LoginActor(c, request, &actorRepo)
	if status < 200 || status > 299 {
		errorMessage = origin.DefaultErrorResponseWithMessage(err.Error(), status)
		c.AbortWithStatusJSON(status, errorMessage)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(actorRepo.Password), []byte(request.Password))
	if err != nil {
		// invalid password
		errorMessage = origin.DefaultErrorResponseWithMessage("invalid username & password", status)
		c.AbortWithStatusJSON(http.StatusUnauthorized, errorMessage)
		return
	}

	//check access
	if actorRepo.Verified != "true" && actorRepo.Active != "true" {
		errorMessage = origin.DefaultErrorResponseWithMessage("account not activated", status)
		c.AbortWithStatusJSON(http.StatusForbidden, errorMessage)
		return
	}

	status, tokenJWTAccess, tokenJWTRefresh, ExpiredRefresh, err = ctr.client.GenerateJWTCustom(c, actorRepo, agent)
	if status < 200 || status > 299 {
		errorMessage = origin.DefaultErrorResponseWithMessage(err.Error(), status)
		c.AbortWithStatusJSON(status, errorMessage)
		return
	}

	status, err = ctr.client.InsertSession(c, tokenJWTRefresh, ExpiredRefresh.Time)
	if status < 200 || status > 299 {
		errorMessage = origin.DefaultErrorResponseWithMessage(err.Error(), status)
		c.AbortWithStatusJSON(status, errorMessage)
		return
	}

	//response
	response = origin.DefaultSuccessResponseWithMessage("login success", status, tokenJWTAccess)

	c.Header("Authorization", "Bearer "+fmt.Sprint(response.Data))
	c.Header("Refresh-Token", "Bearer "+fmt.Sprint(response.Data))
	c.JSON(status, response)
}

//func (h *ControllerAuth) LogoutActor(c *gin.Context) {
//	//req header and del
//	c.Request.Header.Del("Authorization")
//
//	//response
//	c.JSON(http.StatusOK, origin.DefaultSuccessResponseWithMessage("logout success", 200, true))
//}
