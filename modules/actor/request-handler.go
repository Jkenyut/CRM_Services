package actor

import (
	"crm_service/dto"
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
	role, _ := c.Get("role")

	if role != 1 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.DefaultErrorResponseWithMessage("Account Not Authorization", http.StatusUnauthorized))
		return
	}

	// bind to json
	var request RequestActor
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage("required not valid", http.StatusBadRequest))
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
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage("must unsigned number", http.StatusBadRequest))
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
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage("must unsigned number", http.StatusBadRequest))
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
	role, _ := c.Get("role")
	if role != 1 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Not Authorization")
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage("required not valid", http.StatusBadRequest))
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
	role, _ := c.Get("role")
	if role != 1 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Not Authorization")
		return
	}
	actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage("required not valid", http.StatusBadRequest))
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

//}
//
//func (h RequestHandlerActorStruct) ActivateActorById(c *gin.Context) {
//	role, _ := c.Get("role")
//
//	if role != 1 {
//		c.JSON(http.StatusUnauthorized, "Not Authorization")
//		return
//	}
//	actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
//	res, err := h.ctr.ActivateActorById(uint(actorId))
//	if err != nil {
//		if err.Error() == "actor not found" {
//			c.JSON(http.StatusNotFound, "Actor not found")
//			return
//
//		} else if err.Error() == "activate failed" {
//			c.JSON(http.StatusBadRequest, "activate failed")
//			return
//		} else {
//			c.JSON(http.StatusInternalServerError, "Server error")
//			return
//		}
//	}
//	c.JSON(http.StatusOK, res)
//}
//
//func (h RequestHandlerActorStruct) DeactivateActorById(c *gin.Context) {
//	role, _ := c.Get("role")
//	if role != 1 {
//		c.JSON(http.StatusUnauthorized, "Not Authorization")
//		return
//	}
//	actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
//	res, err := h.ctr.DeactivateActorById(uint(actorId))
//	if err != nil {
//		if err.Error() == "actor not found" {
//			c.JSON(http.StatusNotFound, "Actor not found")
//			return
//		} else if err.Error() == "actor is super admin can't deactivate" {
//			c.JSON(http.StatusUnauthorized, "actor is super admin can't deactivate")
//		} else if err.Error() == "deactivate failed" {
//			c.JSON(http.StatusBadRequest, "deactivate failed")
//			return
//		} else {
//			c.JSON(http.StatusInternalServerError, "Server error")
//			return
//		}
//	}
//	c.JSON(http.StatusOK, res)
//}

func (h RequestHandlerActorStruct) LoginActor(c *gin.Context) {

	// get header user-agent
	agent := c.GetHeader("User-Agent")

	//bind to json
	var request RequestActor
	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage("required not valid", http.StatusBadRequest))
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
	c.JSON(http.StatusOK, dto.DefaultSuccessResponseWithMessage("logout success", 200, true))
}
