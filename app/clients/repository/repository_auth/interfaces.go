package repository_auth

import (
	"context"
	"crm_service/app/model/model_actor"
	"crm_service/app/model/origin"
	"github.com/Jkenyut/libs-numeric-go/libs_models/libs_model_jwt"
)

type InterfaceAuth interface {
	LoginActor(ctx context.Context, req model_actor.RequestActor, actorRepository *model_actor.ModelActor) (int, error)
	InsertSession(ctx context.Context, activityId string, agent string, claimRefresh *libs_model_jwt.CustomClaims) (status int, error error)
	CheckSession(ctx context.Context, activityId string) (status int, out origin.JWTModel, error error)
	DeleteSession(ctx context.Context, activityId string) (status int, error error)
}
