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
	"github.com/Jkenyut/libs-numeric-go/libs_tracing"
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
	tr        libs_tracing.InterfaceTracingJaegerOperation
}

func NewControllerAuth(client repository_auth.InterfaceAuth, validate *validator.Validate, conf *config.Config, tr libs_tracing.InterfaceTracingJaegerOperation) InterfaceControllerAuth {
	return &ControllerAuth{
		client:    client,
		validator: validate,
		libsAuth:  libs_auth_jwt.NewClientAuthJWT(conf.JWT),
		tr:        tr,
	}
}

func (ctr *ControllerAuth) LoginActor(c *gin.Context) {
	var response libs_model_response.DefaultResponse

	tr := ctr.tr
	span, _ := tr.SetOperationParent(c.FullPath()+".LoginActor", c)
	tr.TracingTag(c.Request)
	defer span.Finish()

	// get header user-agent
	agent := c.GetHeader("User-Agent")
	tr.SetLog("agent", agent)
	//bind to json
	var request model_actor.RequestActor
	err := c.BindJSON(&request)
	if err != nil {
		tr.SetError("required not valid", err.Error())
		pipeline.AbortWithStatusJSON(c, http.StatusBadRequest, "required not valid")
		return
	}
	tr.SetLog("request", request)

	//logic
	actorRepo, status, errs := ctr.client.LoginActor(tr.OutgoingContext(), request)
	if errs != nil || !helper.IsSuccessStatus(status) {
		tr.SetError("LoginActor", errs.Error())
		pipeline.AbortWithStatusJSON(c, status, errs.Error())
		return
	}
	tr.SetLog("actor", actorRepo)
	tr.SetLog("status", status)

	err = bcrypt.CompareHashAndPassword([]byte(actorRepo.Password), []byte(request.Password))
	if err != nil {
		// invalid password
		tr.SetError("CompareHashAndPassword", err.Error())
		pipeline.AbortWithStatusJSON(c, http.StatusUnauthorized, "invalid username & password")
		return
	}

	//check access
	if actorRepo.Verified != "true" && actorRepo.Active != "true" {
		tr.SetError("account not activated", "account not activated")
		pipeline.AbortWithStatusJSON(c, http.StatusForbidden, "account not activated")
		return
	}

	//newUUID
	newUUID := uuid.New().String()
	externalID := uuid.New().String()
	tr.SetLog("newUUID", newUUID)
	tr.SetLog("externalID", externalID)

	audience := []string{strconv.Itoa(int(actorRepo.RoleID)), agent}
	tr.SetLog("audience", audience)

	tokenJWTAccess, claimsAccess, errs := ctr.libsAuth.GenerateJWTAccessCustom(c, "login", audience, newUUID, externalID, "")
	if errs != nil {
		tr.SetError("GenerateJWTAccessCustom", err.Error())
		pipeline.AbortWithStatusJSON(c, status, errs.Error())
		return
	}
	tr.SetLog("tokenJWTAccess", tokenJWTAccess)
	tr.SetLog("claimsAccess", claimsAccess)

	status, err = ctr.client.InsertSession(c, newUUID, agent, claimsAccess)
	if err != nil || !helper.IsSuccessStatus(status) {
		tr.SetError("InsertSession", err.Error())
		pipeline.AbortWithStatusJSON(c, status, err.Error())
		return
	}
	tr.SetLog("status", status)

	//response
	response = libs_model_response.DefaultSuccessResponseWithMessage("login success", status, tokenJWTAccess)

	c.Header("Authorization", "Bearer "+fmt.Sprint(response.Data))
	tr.SetLog("response", response)
	pipeline.JSON(c, status, "Success Login", response)
}

func (ctr *ControllerAuth) LogoutActor(c *gin.Context) {
	tr := ctr.tr
	span, _ := tr.SetOperationParent("Controller LogoutActor", c)
	tr.TracingTag(c.Request)
	defer span.Finish()

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
