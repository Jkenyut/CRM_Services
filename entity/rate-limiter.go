package entity

import (
	"crm_service/dto"
	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"time"
)

func KeyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func ErrorHandler(c *gin.Context, info ratelimit.Info) {
	c.JSON(429, dto.DefaultErrorResponseWithMessage("Too many requests. Try again in "+time.Until(info.ResetTime).String()))
}
