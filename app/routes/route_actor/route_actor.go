package route_actor

import (
	"crm_service/app/controllers/contoller_actor"
	"crm_service/app/middleware"
	"github.com/gin-gonic/gin"
)

type InterfaceRouteActor interface {
	Handle(router *gin.Engine)
}

type RouterActor struct {
	ctr     contoller_actor.InterfaceControllerActor
	authJWT middleware.InterfacesMiddlewareAuth
}

func NewRouteActor(ctr contoller_actor.InterfaceControllerActor, authJWT middleware.InterfacesMiddlewareAuth) InterfaceRouteActor {
	return &RouterActor{
		ctr:     ctr,
		authJWT: authJWT,
	}

}

func (r *RouterActor) Handle(router *gin.Engine) {
	basePath := "v1/actor"

	actorRouter := router.Group(basePath)
	//
	actorRouter.POST("/register",
		r.ctr.CreateActor,
	)
	//
	//actorRouter.GET("/:id",
	//	r.actorRequestHandler.GetActorById,
	//)
	//actorRouter.GET("",
	//	r.actorRequestHandler.GetAllActor,
	//)
	//
	//actorRouter.PUT("/:id",
	//	r.actorRequestHandler.UpdateActorById,
	//)
	//actorRouter.DELETE("/:id",
	//	r.actorRequestHandler.DeleteActorById,
	//)
	//actorRouter.GET("/:id/activate",
	//	r.actorRequestHandler.ActivateActorById)
	//
	//actorRouter.GET("/:id/deactivate",
	//	r.actorRequestHandler.DeactivateActorById)

	//router.GET("v1/repository-entity_actor/logout",
	//	r.actorRequestHandler.LogoutActor)
}
