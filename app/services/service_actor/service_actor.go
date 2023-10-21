package service_actor

import (
	"crm_service/app/clients/connection"
	"crm_service/app/clients/repository/repository_actor"
	"crm_service/app/config"
	"crm_service/app/controllers/contoller_actor"
	"crm_service/app/routes/route_actor"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func NewServiceActor(router *gin.Engine, conf *config.Config, conn connection.InterfaceConnection, validator *validator.Validate) {
	clientActor := repository_actor.NewClientActor(conf, conn)
	fmt.Println("client")
	controllerActor := contoller_actor.NewControllerActor(clientActor, validator)
	fmt.Print("contoller")
	serviceActor := route_actor.NewRouteActor(controllerActor)
	serviceActor.Handle(router)
}
