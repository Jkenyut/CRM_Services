package services_auth

import (
	"context"
	"crm_service/app/clients/connection"
	"crm_service/app/clients/repository/repository_auth"
	"crm_service/app/config"
	"crm_service/app/controllers/controller_auth/controller_auth_actor"
	"crm_service/app/middleware"
	"crm_service/app/routes/route_auth"
	"github.com/Jkenyut/libs-numeric-go/libs_tracing"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func NewServiceAuth(router *gin.Engine, conf *config.Config, conn connection.InterfaceConnection, validator *validator.Validate) {
	tr := libs_tracing.NewTracingJaegerOperation(context.Background())
	clientAuth := repository_auth.NewClientAuth(conf, conn, tr)
	controllerAuth := controller_auth_actor.NewControllerAuth(clientAuth, validator, conf, tr)
	AuthJWTController := middleware.NewMiddlewareAuth(conf, clientAuth)
	serviceAuth := route_auth.NewRouteAuth(controllerAuth, AuthJWTController)
	serviceAuth.Handle(router)
}
