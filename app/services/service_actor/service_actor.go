package service_actor

import (
	"context"
	"crm_service/app/clients/connection"
	"crm_service/app/clients/repository/repository_actor"
	"crm_service/app/clients/repository/repository_auth"
	"crm_service/app/config"
	"crm_service/app/controllers/controller_actor"
	"crm_service/app/middleware"
	"crm_service/app/routes/route_actor"
	"github.com/Jkenyut/libs-numeric-go/libs_tracing"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func NewServiceActor(router *gin.Engine, conf *config.Config, conn connection.InterfaceConnection, validator *validator.Validate) {
	tr := libs_tracing.NewTracingJaegerOperation(context.Background())
	clientActor := repository_actor.NewClientActor(conf, conn)
	controllerActor := controller_actor.NewControllerActor(clientActor, validator)
	clientAuth := repository_auth.NewClientAuth(conf, conn, tr)
	AuthJWTController := middleware.NewMiddlewareAuth(conf, clientAuth)
	serviceActor := route_actor.NewRouteActor(controllerActor, AuthJWTController)
	serviceActor.Handle(router)
}
