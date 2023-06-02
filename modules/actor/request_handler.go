package actor

import (
	"crm_service/dto"
	"crm_service/repository"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestHandlerActorStruct struct {
	ctr ActorControllerInterface
}

func RequestHandler(
	dbCrud *gorm.DB,
) RequestHandlerActorStruct {
	return RequestHandlerActorStruct{
		ctr: actorControllerStruct{
			actorUseCase: actorUseCaseStruct{
				actorRepository: repository.NewActor(dbCrud),
			},
		}}
}

var validate = validator.New()

func (h RequestHandlerActorStruct) CreateActor(c *gin.Context) {
	request := ActorBody{}
	err := c.Bind(&request)
	fmt.Println(request, err)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	err = validate.Struct(request)

	if err != nil {
		// Validation failed

		for _, err := range err.(validator.ValidationErrors) {
			customErr := fmt.Sprint(err.StructField(), " ", err.ActualTag(), " ", err.Param())
			switch err.Tag() {
			case "required":
				c.JSON(http.StatusUnprocessableEntity, dto.DefaultErrorResponseWithMessage(customErr))
				return
			case "min":
				c.JSON(http.StatusUnprocessableEntity, dto.DefaultErrorResponseWithMessage(customErr))
				return
			case "max":
				c.JSON(http.StatusUnprocessableEntity, dto.DefaultErrorResponseWithMessage(customErr))
				return
			case "alphanum":
				c.JSON(http.StatusUnprocessableEntity, dto.DefaultErrorResponseWithMessage(customErr))
				return
			}
		}
	}
	res, err := h.ctr.CreateActor(request)
	if err != nil {
		if err.Error() == "username already taken" {
			c.JSON(http.StatusConflict, dto.DefaultErrorResponseWithMessage("Username already taken"))
			return
		} else {
			c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage("Server error"))
			return
		}
	}
	c.JSON(http.StatusCreated, res)
}

func (h RequestHandlerActorStruct) GetActorById(c *gin.Context) {
	actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := h.ctr.GetActorById(uint(actorId))
	if err != nil {
		if err.Error() == "actor not found" {
			c.JSON(http.StatusNotFound, dto.DefaultErrorResponseWithMessage("Actor not found"))
			return
		} else {
			c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage("Server error"))
			return
		}

	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerActorStruct) GetAllActor(c *gin.Context) {
	res, err := h.ctr.GetAllActor()
	if err != nil {
		c.JSON(http.StatusNotFound, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerActorStruct) UpdateActorById(c *gin.Context) {
	request := UpdateActorBody{}
	err := c.Bind(&request)
	fmt.Println(request, err)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	err = validate.Struct(request)

	if err != nil {
		// Validation failed

		for _, err := range err.(validator.ValidationErrors) {
			customErr := fmt.Sprint(err.StructField(), " ", err.ActualTag(), " ", err.Param())
			switch err.Tag() {
			case "required":
				c.JSON(http.StatusUnprocessableEntity, dto.DefaultErrorResponseWithMessage(customErr))
				return
			case "min":
				c.JSON(http.StatusUnprocessableEntity, dto.DefaultErrorResponseWithMessage(customErr))
				return
			case "max":
				c.JSON(http.StatusUnprocessableEntity, dto.DefaultErrorResponseWithMessage(customErr))
				return
			case "alphanum":
				c.JSON(http.StatusUnprocessableEntity, dto.DefaultErrorResponseWithMessage(customErr))
				return
			case "eq":
				c.JSON(http.StatusUnprocessableEntity, dto.DefaultErrorResponseWithMessage(customErr))
				return

			}
		}
	}
	res, err := h.ctr.UpdateById(uint(actorId), request)
	if err != nil {
		if err.Error() == "actor not found" {
			c.JSON(http.StatusConflict, dto.DefaultErrorResponseWithMessage("actor not found"))
			return
		} else {
			c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage("Server error"))
			return
		}
	}
	c.JSON(http.StatusCreated, res)
}

func (h RequestHandlerActorStruct) DeleteActorById(c *gin.Context) {
	actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := h.ctr.DeleteActorById(uint(actorId))
	if err != nil {
		if err.Error() == "actor not found" {
			c.JSON(http.StatusNotFound, dto.DefaultErrorResponseWithMessage("Actor not found"))
			return
		} else if err.Error() == "actor is super admin cannot delete" {
			c.JSON(http.StatusUnauthorized, dto.DefaultErrorResponseWithMessage("actor is super admin cannot delete"))
			return
		} else {
			c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage("Server error"))
			return
		}

	}
	c.JSON(http.StatusOK, res)
}
