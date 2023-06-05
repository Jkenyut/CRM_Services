package actor

import (
	"crm_service/dto"
	"crm_service/repository"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"time"

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
	role, _ := c.Get("role")
	if role != 1 {
		c.JSON(http.StatusUnauthorized, dto.DefaultErrorResponseWithMessage("Not Authorization"))
		return
	}
	request := ActorBody{}
	err := c.Bind(&request)

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

	pageStr := c.DefaultQuery("page", "1")
	usernameStr := c.DefaultQuery("username", "")
	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := h.ctr.GetAllActor(uint(page), usernameStr)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerActorStruct) UpdateActorById(c *gin.Context) {
	role, _ := c.Get("role")
	if role != 1 {
		c.JSON(http.StatusUnauthorized, dto.DefaultErrorResponseWithMessage("Not Authorization"))
		return
	}
	request := UpdateActorBody{}
	err := c.Bind(&request)

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
			c.JSON(http.StatusNotFound, dto.DefaultErrorResponseWithMessage("actor not found"))
			return
		} else if err.Error() == "actor is super admin cannot update" {
			c.JSON(http.StatusUnauthorized, dto.DefaultErrorResponseWithMessage("actor is super admin cannot update"))
			return
		} else if err.Error() == "username already taken" {
			c.JSON(http.StatusConflict, dto.DefaultErrorResponseWithMessage("username already taken"))
			return
		} else if err.Error() == "failed to update actor" {
			c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage("failed to update actor"))
			return
		} else {
			c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage("Server error"))
			return
		}
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerActorStruct) DeleteActorById(c *gin.Context) {
	role, _ := c.Get("role")
	if role != 1 {
		c.JSON(http.StatusUnauthorized, dto.DefaultErrorResponseWithMessage("Not Authorization"))
		return
	}
	actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)

	res, err := h.ctr.DeleteActorById(uint(actorId))
	if err != nil {
		if err.Error() == "actor not found" {
			c.JSON(http.StatusNotFound, dto.DefaultErrorResponseWithMessage("Actor not found"))
			return
		} else if err.Error() == "actor is super admin cannot delete" {
			c.JSON(http.StatusUnauthorized, dto.DefaultErrorResponseWithMessage("actor is super admin cannot delete"))
			return
		} else if err.Error() == "failed deleted" {
			c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage("failed deleted"))
			return
		} else {
			c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage("Server error"))
			return
		}

	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerActorStruct) ActivateActorById(c *gin.Context) {
	role, _ := c.Get("role")

	if role != 1 {
		c.JSON(http.StatusUnauthorized, dto.DefaultErrorResponseWithMessage("Not Authorization"))
		return
	}
	actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	res, err := h.ctr.ActivateActorById(uint(actorId))
	if err != nil {
		if err.Error() == "actor not found" {
			c.JSON(http.StatusNotFound, dto.DefaultErrorResponseWithMessage("Actor not found"))
			return

		} else if err.Error() == "activate failed" {
			c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage("activate failed"))
			return
		} else {
			c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage("Server error"))
			return
		}
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerActorStruct) DeactivateActorById(c *gin.Context) {
	role, _ := c.Get("role")
	if role != 1 {
		c.JSON(http.StatusUnauthorized, dto.DefaultErrorResponseWithMessage("Not Authorization"))
		return
	}
	actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	res, err := h.ctr.DeactivateActorById(uint(actorId))
	if err != nil {
		if err.Error() == "actor not found" {
			c.JSON(http.StatusNotFound, dto.DefaultErrorResponseWithMessage("Actor not found"))
			return
		} else if err.Error() == "actor is super admin can't deactivate" {
			c.JSON(http.StatusUnauthorized, dto.DefaultErrorResponseWithMessage("actor is super admin can't deactivate"))
		} else if err.Error() == "deactivate failed" {
			c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage("deactivate failed"))
			return
		} else {
			c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage("Server error"))
			return
		}
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerActorStruct) LoginActor(c *gin.Context) {
	agent := c.GetHeader("User-Agent")
	request := ActorBody{}
	err := c.Bind(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := h.ctr.LoginActor(request, agent)
	c.Header("Authorization", "Bearer "+res.Data)
	if err != nil {
		if err.Error() == "invalid username & password" {
			c.JSON(http.StatusUnauthorized, dto.DefaultErrorResponseWithMessage("invalid username & password"))
			return
		} else if err.Error() == "username not activate" {
			c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage("username not activate"))
			return
		} else if err.Error() == "actor not found" {
			c.JSON(http.StatusNotFound, dto.DefaultErrorResponseWithMessage("actor not found"))
			return
		} else if err.Error() == "failed to generate token" {
			c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage("failed to generate token"))
			return
		} else {
			c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage("Server error"))
			return
		}
	}
	c.JSON(http.StatusOK, res)
}
func (h RequestHandlerActorStruct) LogoutActor(c *gin.Context) {
	start := time.Now()
	c.Request.Header.Del("Authorization")
	c.JSON(http.StatusOK, dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Success logout actor",
		Message:      "Success logout actor",
		ResponseTime: fmt.Sprint(time.Since(start)),
	})
}
