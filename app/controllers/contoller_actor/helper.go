package contoller_actor

import (
	"crm_service/app/model/origin"
	"github.com/gin-gonic/gin"
)

func (ctr *ControllerActor) AbortWithStatusJSON(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, origin.DefaultErrorResponseWithMessage(message, status))
}

func (ctr *ControllerActor) JSON(c *gin.Context, status int, message string, data any) {
	c.JSON(status, origin.DefaultSuccessResponseWithMessage(message, status, data))
}
