package contoller_actor

import (
	"crm_service/app/clients/repository/repository_actor"
	"crm_service/app/model/model_actor"
	"crm_service/app/model/origin"
	"crm_service/app/utils/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

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
	//	c.AbortWithStatusJSON(http.StatusUnauthorized, origin.DefaultErrorResponseWithMessage("Account Not Authorization", http.StatusUnauthorized))
	//	return
	//}

	// bind to json
	var request model_actor.RequestActor
	err := c.Bind(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, origin.DefaultErrorResponseWithMessage("required not valid", http.StatusBadRequest))
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
		c.AbortWithStatusJSON(status, origin.DefaultErrorResponseWithMessage(err.Error(), status))
		return
	}

	var responseActor model_actor.ResponseActor
	responseActor.Username = request.Username
	responseActor.CreatedAt = helper.ConvertTimeToWIB(time.Now())
	responseActor.RoleID = 2
	responseActor.Active = "false"

	c.JSON(status, origin.DefaultSuccessResponseWithMessage("actor created", status, responseActor))
}

func (ctr *ControllerActor) GetActorById(c *gin.Context) {
	var actorRepo model_actor.ModelActor
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, origin.DefaultErrorResponseWithMessage("must unsigned number", http.StatusBadRequest))
		return
	}
	status, err := ctr.client.GetActorById(c, id, &actorRepo)
	//check status
	if status < 200 || status > 299 {
		c.AbortWithStatusJSON(status, origin.DefaultErrorResponseWithMessage(err.Error(), status))
		return
	}

	c.JSON(status, origin.DefaultSuccessResponseWithMessage("actor created", status, actorRepo))
}

//	func (h ControllerActor) GetAllActor(c *gin.Context) {
//		page, err := strconv.ParseUint(c.DefaultQuery("page", "1"), 10, 64)
//		username := c.DefaultQuery("username", "")
//
//		if err != nil {
//			c.AbortWithStatusJSON(http.StatusBadRequest, origin.DefaultErrorResponseWithMessage("must unsigned number", http.StatusBadRequest))
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
		c.AbortWithStatusJSON(http.StatusBadRequest, origin.DefaultErrorResponseWithMessage("required not valid", http.StatusBadRequest))
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

	c.JSON(http.StatusOK, origin.DefaultSuccessResponseWithMessage("Update data actor", status, responseActor))
}

func (ctr *ControllerActor) DeleteActorById(c *gin.Context) {
	envJWT, _ := c.Get("envJWT")
	setJWT := envJWT.(map[string]interface{})

	if setJWT["role"] != "1" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Not Authorization")
		return
	}
	actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, origin.DefaultErrorResponseWithMessage("required not valid", http.StatusBadRequest))
		return
	}

	status, errMessage := ctr.client.DeleteActorById(c, actorId)
	//check status
	if status < 200 || status > 299 {
		c.AbortWithStatusJSON(status, errMessage)
		return
	}
	c.JSON(status, origin.DefaultSuccessResponseWithMessage("success delete data actor", status, "success"))
}

func (ctr *ControllerActor) ActivateActorById(c *gin.Context) {
	//envJWT, _ := c.Get("envJWT")
	//setJWT := envJWT.(map[string]interface{})
	//
	//if setJWT["role"] != "1" {
	//	c.AbortWithStatusJSON(http.StatusUnauthorized, "Not Authorization")
	//	return
	//}
	actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, origin.DefaultErrorResponseWithMessage("required not valid", http.StatusBadRequest))
		return
	}

	status, errMessage := ctr.client.ActivateActorById(c, actorId)
	//check status
	if status < 200 || status > 299 {
		c.AbortWithStatusJSON(status, errMessage)
		return
	}
	c.JSON(status, origin.DefaultSuccessResponseWithMessage("success delete data actor", status, "success"))
}

func (ctr *ControllerActor) DeactivateActorById(c *gin.Context) {
	//envJWT, _ := c.Get("envJWT")
	//setJWT := envJWT.(map[string]interface{})
	//
	//if setJWT["role"] != "1" {
	//	c.AbortWithStatusJSON(http.StatusUnauthorized, "Not Authorization")
	//	return
	//}
	actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, origin.DefaultErrorResponseWithMessage("required not valid", http.StatusBadRequest))
		return
	}

	status, errMessage := ctr.client.DeactivateActorById(c, actorId)
	//check status
	if status < 200 || status > 299 {
		c.AbortWithStatusJSON(status, errMessage)
		return
	}
	c.JSON(status, origin.DefaultSuccessResponseWithMessage("success delete data actor", status, "success"))
}
