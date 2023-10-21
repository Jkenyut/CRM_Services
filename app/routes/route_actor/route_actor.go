package route_actor

import (
	"crm_service/app/controllers/contoller_actor"
	"github.com/gin-gonic/gin"
)

type InterfaceRouteActor interface {
	Handle(router *gin.Engine)
}

type RouterActor struct {
	ctr contoller_actor.InterfacControllerActor
}

func NewRouteActor(ctr contoller_actor.InterfacControllerActor) InterfaceRouteActor {
	return &RouterActor{
		ctr: ctr,
	}

}

func (r *RouterActor) Handle(router *gin.Engine) {
	//basepath := "v1/repository-entity_actor"

	//actorRouter := router.Group(basepath, middleware.Auth)
	//
	//actorRouter.POST("/register",
	//	r.actorRequestHandler.CreateActor,
	//)
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

	router.POST("v1/actor/login",
		r.ctr.LoginActor)

	//router.GET("v1/repository-entity_actor/logout",
	//	r.actorRequestHandler.LogoutActor)
}
