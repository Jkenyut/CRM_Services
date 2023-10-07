package actor

import (
	"crm_service/entity"
	"crm_service/utils/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type RequestHandlerActorStruct struct {
	ctr ControllerActorInterface
}

func RequestHandler(
	dbCrud *gorm.DB,
) RequestHandlerActorStruct {
	return RequestHandlerActorStruct{
		ctr: actorControllerStruct{
			actorRepository: NewActor(dbCrud),
		}}
}

var validate = validator.New()

func (h RequestHandlerActorStruct) CreateActor(c *gin.Context) {
	// get enviroment
	envJWT, _ := c.Get("envJWT")
	setJWT := envJWT.(map[string]interface{})

	if setJWT["role"] != "1" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, entity.DefaultErrorResponseWithMessage("Account Not Authorization", http.StatusUnauthorized))
		return
	}

	// bind to json
	var request RequestActor
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.DefaultErrorResponseWithMessage("required not valid", http.StatusBadRequest))
		return
	}

	//validate
	err = validate.Struct(request)
	if err != nil {
		// Validation failed
		c.AbortWithStatusJSON(helper.ValidateData(err))
		return

	}

	//controller
	res, status, errMessage := h.ctr.CreateActor(c, request)
	//check status
	if status < 200 || status > 299 {
		c.AbortWithStatusJSON(status, errMessage)
		return
	}
	c.JSON(http.StatusCreated, res)
}

func (h RequestHandlerActorStruct) GetActorById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, entity.DefaultErrorResponseWithMessage("must unsigned number", http.StatusBadRequest))
		return
	}
	res, status, errMessage := h.ctr.GetActorById(c, id)
	//check status
	if status < 200 || status > 299 {
		c.AbortWithStatusJSON(status, errMessage)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerActorStruct) GetAllActor(c *gin.Context) {
	page, err := strconv.ParseUint(c.DefaultQuery("page", "1"), 10, 64)
	username := c.DefaultQuery("username", "")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, entity.DefaultErrorResponseWithMessage("must unsigned number", http.StatusBadRequest))
		return
	}

	res, status, errMessage := h.ctr.GetAllActor(c, page, username)
	//check status
	if status < 200 || status > 299 {
		c.AbortWithStatusJSON(status, errMessage)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerActorStruct) UpdateActorById(c *gin.Context) {
	envJWT, _ := c.Get("envJWT")
	setJWT := envJWT.(map[string]interface{})

	if setJWT["role"] != "1" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Not Authorization")
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.DefaultErrorResponseWithMessage("required not valid", http.StatusBadRequest))
		return
	}

	request := RequestUpdateActor{}
	err = c.Bind(&request)
	//validate
	err = validate.Struct(request)
	if err != nil {
		// Validation failed
		c.AbortWithStatusJSON(helper.ValidateData(err))
		return
	}

	res, status, errMessage := h.ctr.UpdateActorById(c, id, request)
	//check status
	if status < 200 || status > 299 {
		c.AbortWithStatusJSON(status, errMessage)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerActorStruct) DeleteActorById(c *gin.Context) {
	envJWT, _ := c.Get("envJWT")
	setJWT := envJWT.(map[string]interface{})

	if setJWT["role"] != "1" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Not Authorization")
		return
	}
	actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.DefaultErrorResponseWithMessage("required not valid", http.StatusBadRequest))
		return
	}

	res, status, errMessage := h.ctr.DeleteActorById(c, actorId)
	//check status
	if status < 200 || status > 299 {
		c.AbortWithStatusJSON(status, errMessage)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerActorStruct) ActivateActorById(c *gin.Context) {
	envJWT, _ := c.Get("envJWT")
	setJWT := envJWT.(map[string]interface{})

	if setJWT["role"] != "1" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Not Authorization")
		return
	}
	actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.DefaultErrorResponseWithMessage("required not valid", http.StatusBadRequest))
		return
	}

	res, status, errMessage := h.ctr.ActivateActorById(c, actorId)
	//check status
	if status < 200 || status > 299 {
		c.AbortWithStatusJSON(status, errMessage)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerActorStruct) DeactivateActorById(c *gin.Context) {
	envJWT, _ := c.Get("envJWT")
	setJWT := envJWT.(map[string]interface{})

	if setJWT["role"] != "1" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Not Authorization")
		return
	}
	actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.DefaultErrorResponseWithMessage("required not valid", http.StatusBadRequest))
		return
	}

	res, status, errMessage := h.ctr.DeactivateActorById(c, actorId)
	//check status
	if status < 200 || status > 299 {
		c.AbortWithStatusJSON(status, errMessage)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerActorStruct) LoginActor(c *gin.Context) {

	// get header user-agent
	agent := c.GetHeader("User-Agent")

	//bind to json
	var request RequestActor
	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, entity.DefaultErrorResponseWithMessage("required not valid", http.StatusBadRequest))
		return
	}

	//controller
	res, status, errMessage := h.ctr.LoginActor(c, request, agent)
	//check status
	if status < 200 || status > 299 {
		c.AbortWithStatusJSON(status, errMessage)
		return
	}

	//response
	c.Header("Authorization", "Bearer "+fmt.Sprint(res.Data))
	c.JSON(status, res)
}

func (h RequestHandlerActorStruct) LogoutActor(c *gin.Context) {
	//req header and del
	c.Request.Header.Del("Authorization")

	//response
	c.JSON(http.StatusOK, entity.DefaultSuccessResponseWithMessage("logout success", 200, true))
}
