package route_auth

import (
	"crm_service/app/controllers/controller_auth/controller_auth_actor"
	"github.com/gin-gonic/gin"
)

type InterfaceRouteAuth interface {
	Handle(router *gin.Engine)
}

type RouterAuth struct {
	ctr controller_auth_actor.InterfaceControllerAuth
}

func NewRouteAuth(ctr controller_auth_actor.InterfaceControllerAuth) InterfaceRouteAuth {
	return &RouterAuth{
		ctr: ctr,
	}
}

func (r *RouterAuth) Handle(router *gin.Engine) {
	router.POST("v1/actor/login",
		r.ctr.LoginActor)

}
