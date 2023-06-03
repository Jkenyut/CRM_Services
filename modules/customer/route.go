package customer

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterCustomerStruct struct {
	customerRequestHandler RequestHandlerCustomerStruct
}

func NewRouter(
	dbCrud *gorm.DB,
) RouterCustomerStruct {
	return RouterCustomerStruct{
		customerRequestHandler: RequestHandler(
			dbCrud,
		),
	}
}

func (r RouterCustomerStruct) Handle(router *gin.Engine) {
	basepath := "v1/customer"
	customerRouter := router.Group(basepath)

	customerRouter.POST("/register",
		r.customerRequestHandler.CreateCustomer,
	)

	customerRouter.GET("/:id",
		r.customerRequestHandler.GetCustomerById,
	)
	customerRouter.GET("",
		r.customerRequestHandler.GetAllCustomer,
	)

	customerRouter.PUT("/:id",
		r.customerRequestHandler.UpdateCustomerById,
	)
	customerRouter.DELETE("/:id",
		r.customerRequestHandler.DeleteCustomerById,
	)
}
