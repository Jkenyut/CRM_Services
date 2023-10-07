package model

import (
	"crm_service/entity"
	"fmt"
	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"net/http"
)

func KeyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func ErrorHandler(c *gin.Context, info ratelimit.Info) {
	fmt.Print(info)
	c.AbortWithStatusJSON(http.StatusTooManyRequests, entity.DefaultErrorResponseWithMessage("error", http.StatusTooManyRequests))
}
