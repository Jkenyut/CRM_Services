package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterUser struct {
	UserRequestHandler RequestHandlerUser
}

func NewRouter(
	dbCrud *gorm.DB,
) RouterUser {
	return RouterUser{
		UserRequestHandler: NewUserRequestHandler(
			dbCrud,
		),
	}
}

func (r RouterUser) Handle(router *gin.Engine) {
	basepath := "/user"
	user := router.Group(basepath)

	user.POST("/register",
		r.UserRequestHandler.CreateUser,
	)

	user.GET("/register/:id",
		r.UserRequestHandler.GetUserById,
	)
}
