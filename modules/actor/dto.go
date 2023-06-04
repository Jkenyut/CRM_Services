package actor

import (
	"crm_service/dto"
	"crm_service/entity"
	"github.com/dgrijalva/jwt-go"
)

type ActorBody struct {
	Username string `json:"username" validate:"required,min=1,max=100,alphanum"`
	Password string `json:"password,omitempty" validate:"required,min=6,max=100"`
}

type UpdateActorBody struct {
	Username string `json:"username" validate:"min=1,max=100,alphanum"`
	Password string `json:"password,omitempty" validate:"min=6,max=100"`
	Verified string `json:"verified" validate:"eq=true|eq=false"`
	Active   string `json:"active" validate:"eq=true|eq=false"`
}

type SuccessCreate struct {
	dto.ResponseMeta
	Data ActorBody `json:"data"`
}

type FindActor struct {
	dto.ResponseMeta
	Data entity.Actor `json:"data"`
}

type FindAllActor struct {
	dto.ResponseMeta
	Page       uint           `json:"page,omitempty"`
	PerPage    uint           `json:"per_page,omitempty"`
	Total      int            `json:"total,omitempty"`
	TotalPages uint           `json:"total_pages,omitempty"`
	Data       []entity.Actor `json:"data"`
}

type CustomClaims struct {
	Role      uint   `json:"role"`
	UserAgent string `json:"user_agent"`
	jwt.StandardClaims
}

type SuccessLogin struct {
	dto.ResponseMeta
	Data string `json:"data"`
}
