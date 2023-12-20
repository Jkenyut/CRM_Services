package controller_actor

import "github.com/gin-gonic/gin"

type InterfaceControllerActor interface {
	CreateActor(c *gin.Context)
	GetActorById(c *gin.Context)
	GetAllActor(c *gin.Context)
	UpdateActorById(c *gin.Context)
	DeleteActorById(c *gin.Context)
}
