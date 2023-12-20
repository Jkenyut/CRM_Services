package controller_auth_actor

import "github.com/gin-gonic/gin"

type InterfaceControllerAuth interface {
	LoginActor(c *gin.Context)
	LogoutActor(c *gin.Context)
}
