package actor

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
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

//var validate = validator.New()

//func (h RequestHandlerActorStruct) CreateActor(c *gin.Context) {
//	role, _ := c.Get("role")
//	if role != 1 {
//		c.JSON(http.StatusUnauthorized, "Not Authorizationwinc")
//		return
//	}
//	request := ActorRequest{}
//	err := c.Bind(&request)
//
//	if err != nil {
//		c.JSON(http.StatusBadRequest, "w")
//		return
//	}
//	err = validate.Struct(request)
//
//	if err != nil {
//		// Validation failed
//
//		for _, err := range err.(validator.ValidationErrors) {
//			customErr := fmt.Sprint(err.StructField(), " ", err.ActualTag(), " ", err.Param())
//			switch err.Tag() {
//			case "required":
//				c.JSON(http.StatusUnprocessableEntity, customErr)
//				return
//			case "min":
//				c.JSON(http.StatusUnprocessableEntity, customErr)
//				return
//			case "max":
//				c.JSON(http.StatusUnprocessableEntity, customErr)
//				return
//			case "alphanum":
//				c.JSON(http.StatusUnprocessableEntity, customErr)
//				return
//			}
//		}
//	}
//	res, err := h.ctr.CreateActor(request)
//	if err != nil {
//		if err.Error() == "username already taken" {
//			c.JSON(http.StatusConflict, "Username already taken")
//			return
//		} else {
//			c.JSON(http.StatusInternalServerError, "Server error")
//			return
//		}
//	}
//	c.JSON(http.StatusCreated, res)
//}
//
//func (h RequestHandlerActorStruct) GetActorById(c *gin.Context) {
//	actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, "wsj cf")
//		return
//	}
//
//	res, err := h.ctr.GetActorById(uint(actorId))
//	if err != nil {
//		if err.Error() == "actor not found" {
//			c.JSON(http.StatusNotFound, "Actor not found")
//			return
//		} else {
//			c.JSON(http.StatusInternalServerError, "Server error")
//			return
//		}
//
//	}
//	c.JSON(http.StatusOK, res)
//}
//
//func (h RequestHandlerActorStruct) GetAllActor(c *gin.Context) {
//
//	pageStr := c.DefaultQuery("page", "1")
//	usernameStr := c.DefaultQuery("username", "")
//	page, err := strconv.ParseUint(pageStr, 10, 64)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, "wjf ")
//		return
//	}
//
//	res, err := h.ctr.GetAllActor(uint(page), usernameStr)
//	if err != nil {
//		c.JSON(http.StatusNotFound, err.Error())
//		return
//	}
//	c.JSON(http.StatusOK, res)
//}
//
//func (h RequestHandlerActorStruct) UpdateActorById(c *gin.Context) {
//	role, _ := c.Get("role")
//	if role != 1 {
//		c.JSON(http.StatusUnauthorized, "Not Authorization")
//		return
//	}
//	request := UpdateActorRequest{}
//	err := c.Bind(&request)
//
//	if err != nil {
//		c.JSON(http.StatusBadRequest, "weknfw")
//		return
//	}
//
//	actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, "wjw")
//		return
//	}
//
//	err = validate.Struct(request)
//
//	if err != nil {
//		// Validation failed
//
//		for _, err := range err.(validator.ValidationErrors) {
//			customErr := fmt.Sprint(err.StructField(), " ", err.ActualTag(), " ", err.Param())
//			switch err.Tag() {
//			case "required":
//				c.JSON(http.StatusUnprocessableEntity, customErr)
//				return
//			case "min":
//				c.JSON(http.StatusUnprocessableEntity, customErr)
//				return
//			case "max":
//				c.JSON(http.StatusUnprocessableEntity, customErr)
//				return
//			case "alphanum":
//				c.JSON(http.StatusUnprocessableEntity, customErr)
//				return
//			case "eq":
//				c.JSON(http.StatusUnprocessableEntity, customErr)
//				return
//
//			}
//		}
//	}
//	res, err := h.ctr.UpdateById(uint(actorId), request)
//	if err != nil {
//		if err.Error() == "actor not found" {
//			c.JSON(http.StatusNotFound, "actor not found")
//			return
//		} else if err.Error() == "actor is super admin cannot update" {
//			c.JSON(http.StatusUnauthorized, "actor is super admin cannot update")
//			return
//		} else if err.Error() == "username already taken" {
//			c.JSON(http.StatusConflict, "username already taken")
//			return
//		} else if err.Error() == "failed to update actor" {
//			c.JSON(http.StatusBadRequest, "failed to update actor")
//			return
//		} else {
//			c.JSON(http.StatusInternalServerError, "Server error")
//			return
//		}
//	}
//	c.JSON(http.StatusOK, res)
//}
//
//func (h RequestHandlerActorStruct) DeleteActorById(c *gin.Context) {
//	role, _ := c.Get("role")
//	if role != 1 {
//		c.JSON(http.StatusUnauthorized, "Not Authorization")
//		return
//	}
//	actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
//
//	res, err := h.ctr.DeleteActorById(uint(actorId))
//	if err != nil {
//		if err.Error() == "actor not found" {
//			c.JSON(http.StatusNotFound, "Actor not found")
//			return
//		} else if err.Error() == "actor is super admin cannot delete" {
//			c.JSON(http.StatusUnauthorized, "actor is super admin cannot delete")
//			return
//		} else if err.Error() == "failed deleted" {
//			c.JSON(http.StatusBadRequest, "failed deleted")
//			return
//		} else {
//			c.JSON(http.StatusInternalServerError, "Server error")
//			return
//		}
//
//	}
//	c.JSON(http.StatusOK, res)
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
	agent := c.GetHeader("User-Agent")
	var request RequestActor
	err := c.Bind(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, "error")
	}

	res, status, errMessage := h.ctr.LoginActor(c, request, agent)

	if status < 200 || status > 299 {
		c.AbortWithStatusJSON(status, errMessage)
		return
	}
	c.Header("Authorization", "Bearer "+fmt.Sprint(res.Data))
	c.JSON(status, res)
}

//func (h RequestHandlerActorStruct) LogoutActor(c *gin.Context) {
//	start := time.Now()
//	c.RequestActor.Header.Del("Authorization")
//	c.JSON(http.StatusOK, dto.ResponseMeta{
//		Success:      true,
//		Message:      "Success logout actor",
//		ResponseTime: fmt.Sprint(time.Since(start)),
//	})
//}
