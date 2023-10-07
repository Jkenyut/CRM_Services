package actor

import (
	"crm_service/model"
)

type RequestActor struct {
	Username string `json:"username,omitempty" validate:"required,min=1,max=100,alphanum"`
	Password string `json:"password,omitempty" validate:"required,min=6,max=100"`
}

type RequestApproval struct {
	ID uint64 `json:"adminID,omitempty"`
}

type RequestUpdateActor struct {
	Username string `json:"username,omitempty" validate:"min=1,max=100,alphanum"`
	Verified string `json:"verified,omitempty" validate:"eq=true|eq=false"`
	Active   string `json:"active,omitempty" validate:"eq=true|eq=false"`
}

type FindAllActor struct {
	Page       uint64        `json:"page,omitempty"`
	PerPage    uint64        `json:"per_page,omitempty"`
	TotalPages uint64        `json:"total_pages,omitempty"`
	Data       []model.Actor `json:"data,omitempty"`
}

type customClaimsJWT struct {
	Role      string `json:"role,omitempty"`
	UserAgent string `json:"user_agent,omitempty"`
}
