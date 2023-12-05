package middleware

import (
	"crm_service/app/model/origin"
	"github.com/gin-gonic/gin"
)

func (m *AuthMiddleware) AbortWithStatusJSON(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, origin.DefaultErrorResponseWithMessage(message, status))
}

func (m *AuthMiddleware) JSON(c *gin.Context, status int, message string, data any) {
	c.JSON(status, origin.DefaultSuccessResponseWithMessage(message, status, data))
}
