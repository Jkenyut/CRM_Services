package services_auth

import (
	"crm_service/app/clients/connection"
	"crm_service/app/clients/repository/repository_auth"
	"crm_service/app/config"
	"crm_service/app/controllers/controller_auth/controller_auth_actor"
	"crm_service/app/routes/route_auth"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func NewServiceAuth(router *gin.Engine, conf *config.Config, conn connection.InterfaceConnection, validator *validator.Validate) {
	clientAuth := repository_auth.NewClientAuth(conf, conn)
	controllerAuth := controller_auth_actor.NewControllerAuth(clientAuth, validator)
	serviceAuth := route_auth.NewRouteAuth(controllerAuth)
	serviceAuth.Handle(router)
}
