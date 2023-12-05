package pipeline

import (
	"crm_service/app/model/origin"
	"github.com/gin-gonic/gin"
)

type InterfacePipeline interface {
	AbortWithStatusJSON(status int, message string)
	JSON(status int, message string, data any)
}

type Pipeline struct {
	c *gin.Context
}

func NewPipeline(c *gin.Context) InterfacePipeline {
	return &Pipeline{c: c}
}

func (r *Pipeline) AbortWithStatusJSON(status int, message string) {
	r.c.AbortWithStatusJSON(status, origin.DefaultErrorResponseWithMessage(message, status))
}

func (r *Pipeline) JSON(status int, message string, data any) {
	r.c.JSON(status, origin.DefaultSuccessResponseWithMessage(message, status, data))
}

func AbortWithStatusJSON(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, origin.DefaultErrorResponseWithMessage(message, status))
}

func JSON(c *gin.Context, status int, message string, data any) {
	c.JSON(status, origin.DefaultSuccessResponseWithMessage(message, status, data))
}
