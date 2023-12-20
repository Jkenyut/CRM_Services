package route_actor

import (
	"crm_service/app/controllers/controller_actor"
	"crm_service/app/middleware"
	"github.com/gin-gonic/gin"
)

type InterfaceRouteActor interface {
	Handle(router *gin.Engine)
}

type RouterActor struct {
	ctr     controller_actor.InterfaceControllerActor
	authJWT middleware.InterfacesMiddlewareAuth
}

func NewRouteActor(ctr controller_actor.InterfaceControllerActor, authJWT middleware.InterfacesMiddlewareAuth) InterfaceRouteActor {
	return &RouterActor{
		ctr:     ctr,
		authJWT: authJWT,
	}
}

func (r *RouterActor) Handle(router *gin.Engine) {
	basePath := "v1/actor"
	actorRouter := router.Group(basePath)
	//
	actorRouter.POST("", r.authJWT.Auth,
		r.ctr.CreateActor,
	)
	//
	actorRouter.GET("/:id", r.authJWT.Auth,
		r.ctr.GetActorById,
	)
	actorRouter.GET("", r.authJWT.Auth,
		r.ctr.GetAllActor,
	)

	actorRouter.PUT("/:id", r.authJWT.Auth,
		r.ctr.UpdateActorById,
	)

	actorRouter.DELETE("/:id", r.authJWT.Auth,
		r.ctr.DeleteActorById,
	)
}
