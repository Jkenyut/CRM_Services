package controller_auth_actor

import (
	"crm_service/app/clients/repository/repository_auth"
	"crm_service/app/config"
	"crm_service/app/middleware/pipeline"
	"crm_service/app/model/model_actor"
	"crm_service/app/utils/helper"
	"encoding/json"
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

type Comment struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	PostID  string `json:"postId"`
}

func (ctr *ControllerAuth) LoginActor(c *gin.Context) {
	var response libs_model_response.DefaultResponse
	init := libs_tracing.NewTracingJaeger("libs-numeric-crm")
	_, lo := init.InitJaeger()
	defer lo.Close()
	tr := ctr.tr

	SpanParent, _ := tr.SetOperationParent(c, c.Request, "LoginActor")
	defer SpanParent.Finish()

	// Inject the SpanParent context into the HTTP request headers
	// Set some tags to the span
	url := fmt.Sprintf("http://localhost:9090/api/v1/comments?postId=%s", "1-ab-2")
	req, err := http.NewRequestWithContext(c, "GET", url, nil)

	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to send request: %v", err))
		return
	}

	err = libs_tracing.Inject(SpanParent, req)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to send request: %v", err))
		return
	}
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	// Baca respons dari layanan lain
	// Di sini Anda dapat memproses respons sesuai kebutuhan
	// Contoh: mengembalikan konten respons kepada klien Gin
	// Decode the JSON response
	var data []Comment
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	tr.SetLog("log", data)

	// get header user-agent
	agent := c.GetHeader("User-Agent")
	tr.SetLog("agent", agent)
	//bind to json
	var request model_actor.RequestActor
	err = c.BindJSON(&request)
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
	span, _ := tr.SetOperationParent(c, c.Request, "Controller LogoutActor")
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
