package services_customer

import (
	"crm_service/app/clients/connection"
	"crm_service/app/clients/repository/repository_auth"
	"crm_service/app/clients/repository/repository_customer"
	"crm_service/app/config"
	"crm_service/app/controllers/controller_customer"
	"crm_service/app/middleware"
	"crm_service/app/routes/route_customer"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func NewServiceCustomer(router *gin.Engine, conf *config.Config, conn connection.InterfaceConnection, validator *validator.Validate) {
	clientAuth := repository_auth.NewClientAuth(conf, conn)
	AuthJWTController := middleware.NewMiddlewareAuth(conf, clientAuth)
	clientCustomer := repository_customer.NewClientCustomer(conf, conn)
	controllerCustomer := controller_customer.NewControllerCustomer(clientCustomer, validator)
	serviceAuth := route_customer.NewRouteCustomer(controllerCustomer, AuthJWTController)
	serviceAuth.Handle(router)
}
