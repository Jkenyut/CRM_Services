package pipeline

import (
	"crm_service/app/model/origin"
	"github.com/gin-gonic/gin"
)

func AbortWithStatusJSON(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, origin.DefaultErrorResponseWithMessage(message, status))
}

func JSON(c *gin.Context, status int, message string, data any) {
	c.JSON(status, origin.DefaultSuccessResponseWithMessage(message, status, data))
}
