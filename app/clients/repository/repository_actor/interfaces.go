package repository_actor

import (
	"context"
	"crm_service/app/model/model_actor"
)

type InterfaceRepositoryActor interface {
	CreateActor(ctx context.Context, req *model_actor.RequestActor) (int, error)
	//CreateApproval(ctx context.Context, req *entity_actor.RequestApproval) (int, error)
	//GetActorByUsername(ctx context.Context, req entity_actor.RequestActor, actorRepository *ModelActor) (int, error)
	GetActorById(ctx context.Context, id uint64, actorRepository *model_actor.ModelActor) (int, error)
	//GetAllActor(ctx context.Context, page uint64, limit uint64, username string, actorRepository *[]ModelActor) (int, error)
	//GetCountRowsActor(ctx context.Context, actorRepository *ModelActor) (int, error)
	UpdateActorById(ctx context.Context, id uint64, updateActor model_actor.RequestUpdateActor) (int, error)
	//DeleteActorById(ctx context.Context, id uint64) (int, error)
	//ActivateActorById(ctx context.Context, id uint64) (int, error)
	//DeactivateActorById(ctx context.Context, id uint64) (int, error)
	LoginActor(ctx context.Context, req model_actor.RequestActor, actorRepository *model_actor.ModelActor) (int, error)
	//GenerateJWT(ctx context.Context, req model_actor.ModelActor, agent string) (string, string, error)
}
