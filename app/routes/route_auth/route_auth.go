package route_auth

import (
	"crm_service/app/controllers/controller_auth/controller_auth_actor"
	"crm_service/app/middleware"
	"github.com/gin-gonic/gin"
)

type InterfaceRouteAuth interface {
	Handle(router *gin.Engine)
}

type RouterAuth struct {
	ctr     controller_auth_actor.InterfaceControllerAuth
	authJWT middleware.InterfacesMiddlewareAuth
}

func NewRouteAuth(ctr controller_auth_actor.InterfaceControllerAuth, authJWT middleware.InterfacesMiddlewareAuth) InterfaceRouteAuth {
	return &RouterAuth{
		ctr:     ctr,
		authJWT: authJWT,
	}
}

func (r *RouterAuth) Handle(router *gin.Engine) {
	router.POST("v1/actor/login",
		r.ctr.LoginActor)
	router.POST("v1/actor/logout", r.authJWT.Auth,
		r.ctr.LogoutActor)

}
