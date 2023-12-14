package controller_auth_actor

import (
	"crm_service/app/clients/repository/repository_auth"
	"crm_service/app/config"
	"crm_service/app/model/model_actor"
	"fmt"
	"github.com/Jkenyut/libs-numeric-go/libs_auth/libs_auth_jwt"
	"github.com/Jkenyut/libs-numeric-go/libs_models/libs_model_response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type ControllerAuth struct {
	client    repository_auth.InterfaceAuth
	validator *validator.Validate
	libs_auth libs_auth_jwt.InterfacesAuthJWT
}

func NewControllerAuth(client repository_auth.InterfaceAuth, validate *validator.Validate, conf *config.Config) InterfaceControllerAuth {
	return &ControllerAuth{
		client:    client,
		validator: validate,
		libs_auth: libs_auth_jwt.NewClientAuthJWT(conf.JWT),
	}
}

func (ctr *ControllerAuth) LoginActor(c *gin.Context) {
	var actorRepo model_actor.ModelActor
	var response libs_model_response.DefaultResponse
	var errorMessage libs_model_response.DefaultResponse
	var status int

	// get header user-agent
	agent := c.GetHeader("User-Agent")

	//bind to json
	var request model_actor.RequestActor
	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, libs_model_response.DefaultErrorResponseWithMessage("required not valid", http.StatusBadRequest))
		return
	}

	//logic
	status, err = ctr.client.LoginActor(c, request, &actorRepo)
	if status < 200 || status > 299 {
		errorMessage = libs_model_response.DefaultErrorResponseWithMessage(err.Error(), status)
		c.AbortWithStatusJSON(status, errorMessage)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(actorRepo.Password), []byte(request.Password))
	if err != nil {
		// invalid password
		errorMessage = libs_model_response.DefaultErrorResponseWithMessage("invalid username & password", status)
		c.AbortWithStatusJSON(http.StatusUnauthorized, errorMessage)
		return
	}

	//check access
	if actorRepo.Verified != "true" && actorRepo.Active != "true" {
		errorMessage = libs_model_response.DefaultErrorResponseWithMessage("account not activated", status)
		c.AbortWithStatusJSON(http.StatusForbidden, errorMessage)
		return
	}

	//newUUID
	newUUID := uuid.New().String()
	externalID := uuid.New().String()

	audience := []string{strconv.Itoa(int(actorRepo.RoleID)), agent}

	tokenJWTAccess, claimsAccess, err := ctr.libs_auth.GenerateJWTAccessCustom(c, "login", audience, newUUID, externalID, "")
	if err != nil {
		errorMessage = libs_model_response.DefaultErrorResponseWithMessage(err.Error(), status)
		c.AbortWithStatusJSON(status, errorMessage)
		return
	}

	//status, tokenJWTRefresh, ExpiredRefresh, err = ctr.client.GenerateJWTRefreshCustom(c, strconv.Itoa(int(actorRepo.RoleID)), agent, newUUID)
	//if status < 200 || status > 299 {
	//	errorMessage = libs_model_response.DefaultErrorResponseWithMessage()(err.Error(), status)
	//	c.AbortWithStatusJSON(status, errorMessage)
	//	return
	//}

	status, err = ctr.client.InsertSession(c, newUUID, agent, claimsAccess)
	if status < 200 || status > 299 {
		errorMessage = libs_model_response.DefaultErrorResponseWithMessage(err.Error(), status)
		c.AbortWithStatusJSON(status, errorMessage)
		return
	}

	//response
	response = libs_model_response.DefaultSuccessResponseWithMessage("login success", status, tokenJWTAccess)

	c.Header("Authorization", "Bearer "+fmt.Sprint(response.Data))
	c.JSON(status, response)
}

//func (h *ControllerAuth) LogoutActor(c *gin.Context) {
//	//req header and del
//	c.Request.Header.Del("Authorization")
//
//	//response
//	c.JSON(http.StatusOK, origin.DefaultSuccessResponseWithMessage("logout success", 200, true))
//}
