package contoller_actor

import "github.com/gin-gonic/gin"

type InterfaceControllerActor interface {
	CreateActor(c *gin.Context)
}
