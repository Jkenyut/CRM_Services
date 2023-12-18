package controller_auth_actor

import (
	"crm_service/app/clients/repository/repository_auth"
	"crm_service/app/config"
	"crm_service/app/middleware/pipeline"
	"crm_service/app/model/model_actor"
	"crm_service/app/utils/helper"
	"fmt"
	"github.com/Jkenyut/libs-numeric-go/libs_auth/libs_auth_jwt"
	"github.com/Jkenyut/libs-numeric-go/libs_models/libs_model_jwt"
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
	libsAuth  libs_auth_jwt.InterfacesAuthJWT
}

func NewControllerAuth(client repository_auth.InterfaceAuth, validate *validator.Validate, conf *config.Config) InterfaceControllerAuth {
	return &ControllerAuth{
		client:    client,
		validator: validate,
		libsAuth:  libs_auth_jwt.NewClientAuthJWT(conf.JWT),
	}
}

func (ctr *ControllerAuth) LoginActor(c *gin.Context) {
	var actorRepo model_actor.ModelActor
	var response libs_model_response.DefaultResponse
	var status int

	// get header user-agent
	agent := c.GetHeader("User-Agent")

	//bind to json
	var request model_actor.RequestActor
	err := c.BindJSON(&request)
	if err != nil {
		pipeline.AbortWithStatusJSON(c, http.StatusBadRequest, "required not valid")
		return
	}

	//logic
	status, err = ctr.client.LoginActor(c, request, &actorRepo)
	if err != nil || !helper.IsSuccessStatus(status) {
		pipeline.AbortWithStatusJSON(c, status, err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(actorRepo.Password), []byte(request.Password))
	if err != nil {
		// invalid password
		pipeline.AbortWithStatusJSON(c, http.StatusUnauthorized, "invalid username & password")
		return
	}

	//check access
	if actorRepo.Verified != "true" && actorRepo.Active != "true" {
		pipeline.AbortWithStatusJSON(c, http.StatusForbidden, "account not activated")
		return
	}

	//newUUID
	newUUID := uuid.New().String()
	externalID := uuid.New().String()

	audience := []string{strconv.Itoa(int(actorRepo.RoleID)), agent}

	tokenJWTAccess, claimsAccess, err := ctr.libsAuth.GenerateJWTAccessCustom(c, "login", audience, newUUID, externalID, "")
	if err != nil {
		pipeline.AbortWithStatusJSON(c, status, err.Error())
		return
	}

	status, err = ctr.client.InsertSession(c, newUUID, agent, claimsAccess)
	if err != nil || !helper.IsSuccessStatus(status) {
		pipeline.AbortWithStatusJSON(c, status, err.Error())
		return
	}

	//response
	response = libs_model_response.DefaultSuccessResponseWithMessage("login success", status, tokenJWTAccess)

	c.Header("Authorization", "Bearer "+fmt.Sprint(response.Data))
	pipeline.JSON(c, status, "Success Login", response)
}

func (ctr *ControllerAuth) LogoutActor(c *gin.Context) {
	envJWT, ok := c.Get("envJWT")
	if !ok {
		pipeline.AbortWithStatusJSON(c, http.StatusForbidden, "env jwt not found")
		return
	}
	setJWT := envJWT.(*libs_model_jwt.CustomClaims)
	//req header and del

	status, err := ctr.client.DeleteSession(c, setJWT.Subject)
	fmt.Println(setJWT)
	if err != nil || !helper.IsSuccessStatus(status) {
		pipeline.AbortWithStatusJSON(c, status, err.Error())
		return
	}
	c.Request.Header.Del("Authorization")
	//response
	c.JSON(http.StatusOK, libs_model_response.DefaultSuccessResponseWithMessage("logout success", 200, "Success Logout"))
}
