package repository_auth

import (
	"context"
	"crm_service/app/model/model_actor"
	"crm_service/app/model/origin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type InterfaceAuth interface {
	LoginActor(ctx context.Context, req model_actor.RequestActor, actorRepository *model_actor.ModelActor) (int, error)
	InsertSession(ctx context.Context, activity_id string, tokenRefresh string, expiredRefresh time.Time) (status int, error error)
	CheckSession(ctx context.Context, activity_id string) (status int, out origin.JWTModel, error error)
	GenerateJWTAccessCustom(ctx context.Context, req int, agent string, activityId string) (int, string, error)
	GenerateJWTRefreshCustom(ctx context.Context, req int, agent string, activityId string) (int, string, *jwt.NumericDate, error)
}
