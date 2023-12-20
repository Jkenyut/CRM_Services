package route_customer

import (
	"crm_service/app/controllers/controller_customer"
	"crm_service/app/middleware"
	"github.com/gin-gonic/gin"
)

type InterfaceRouteCustomer interface {
	Handle(router *gin.Engine)
}

type RouterCustomer struct {
	ctr     controller_customer.InterfaceControllerCustomer
	authJWT middleware.InterfacesMiddlewareAuth
}

func NewRouteCustomer(ctr controller_customer.InterfaceControllerCustomer, authJWT middleware.InterfacesMiddlewareAuth) InterfaceRouteCustomer {
	return &RouterCustomer{
		ctr:     ctr,
		authJWT: authJWT,
	}
}

func (r *RouterCustomer) Handle(router *gin.Engine) {
	basepath := "v1/customer"
	customerRouter := router.Group(basepath)

	customerRouter.POST("",
		r.ctr.CreateCustomer,
	)
	customerRouter.GET("/email", r.ctr.GetCustomerByEmail)
	customerRouter.GET("/:id", r.ctr.GetCustomerById)
	customerRouter.GET("", r.ctr.GetAllCustomer)
	customerRouter.PUT("/:id", r.ctr.UpdateCustomerById)
	customerRouter.DELETE("/:id", r.ctr.DeleteCustomerById)

}
