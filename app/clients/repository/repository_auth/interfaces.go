package repository_auth

import (
	"context"
	"crm_service/app/model/model_actor"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type InterfaceAuth interface {
	LoginActor(ctx context.Context, req model_actor.RequestActor, actorRepository *model_actor.ModelActor) (int, error)
	InsertSession(ctx context.Context, tokenRefresh string, expiredRefresh time.Time) (status int, error error)
	GenerateJWTCustom(ctx context.Context, req model_actor.ModelActor, agent string) (int, string, string, *jwt.NumericDate, error)
}
