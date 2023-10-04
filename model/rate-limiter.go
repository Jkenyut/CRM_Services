package model

import (
	"crm_service/dto"
	"fmt"
	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
)

func KeyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func ErrorHandler(c *gin.Context, info ratelimit.Info) {
	fmt.Print(info)
	c.JSON(429, dto.DefaultErrorResponseWithMessage("error", "1892", "400"))
}
