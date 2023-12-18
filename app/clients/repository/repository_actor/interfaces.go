package repository_actor

import (
	"context"
	"crm_service/app/model/model_actor"
)

type InterfaceRepositoryActor interface {
	CreateActor(ctx context.Context, req *model_actor.RequestActor) (int, error)
	GetActorByUsername(ctx context.Context, req *model_actor.RequestActor, actorRepository *model_actor.ModelActor) (int, error)
	GetActorById(ctx context.Context, ID uint64) (status int, err error, actorRepository model_actor.ModelActor)
	GetAllActor(ctx context.Context, page uint64, limit uint64, username string) (status int, err error, actorRepository []model_actor.ModelActor)
	GetCountRowsActor(ctx context.Context) (status int, err error, actorRepository model_actor.ModelActor)
	UpdateActorById(ctx context.Context, id uint64, updateActor model_actor.RequestUpdateActor) (int, error)
	DeleteActorById(ctx context.Context, id uint64) (int, error)
	ActivateActorById(ctx context.Context, id uint64) (int, error)
	DeactivateActorById(ctx context.Context, id uint64) (int, error)
}
