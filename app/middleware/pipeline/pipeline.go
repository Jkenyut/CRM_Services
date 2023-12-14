package pipeline

import (
	"github.com/Jkenyut/libs-numeric-go/libs_models/libs_model_response"
	"github.com/gin-gonic/gin"
)

func AbortWithStatusJSON(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, libs_model_response.DefaultErrorResponseWithMessage(message, status))
}

func JSON(c *gin.Context, status int, message string, data any) {
	c.JSON(status, libs_model_response.DefaultSuccessResponseWithMessage(message, status, data))
}
