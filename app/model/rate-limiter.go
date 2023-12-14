package model

import (
	"fmt"
	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/Jkenyut/libs-numeric-go/libs_models/libs_model_response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func KeyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func ErrorHandler(c *gin.Context, info ratelimit.Info) {
	fmt.Print(info)
	c.AbortWithStatusJSON(http.StatusTooManyRequests, libs_model_response.DefaultErrorResponseWithMessage("error", http.StatusTooManyRequests))
}
