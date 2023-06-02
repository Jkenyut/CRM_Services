package actor

import (
	"crm_service/dto"
	"crm_service/entity"
)

type ActorBody struct {
	Username string `json:"username" validate:"required,min=1,max=100,alphanum"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type UpdateActorBody struct {
	Username string `json:"username" validate:"required,min=1,max=100,alphanum"`
	Password string `json:"password" validate:"required,min=6,max=100"`
	Verified string `json:"verified" validate:"eq=true|eq=false"`
	Active   string `json:"active" validate:"eq=true|eq=false"`
}

type ResponseActorBody struct {
	Username string `json:"username"`
}

type SuccessCreate struct {
	dto.ResponseMeta
	Data ResponseActorBody `json:"data"`
}

type FindActor struct {
	dto.ResponseMeta
	Data entity.Actor `json:"data"`
}

type FindAllActor struct {
	dto.ResponseMeta
	Data []entity.Actor `json:"data"`
}
