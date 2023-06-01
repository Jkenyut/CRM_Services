package user

import (
	"crud/dto"
	"crud/entity"
)

type UserParam struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SuccessCreate struct {
	dto.ResponseMeta
	Data UserParam `json:"data"`
}

type FindUser struct {
	dto.ResponseMeta
	Data entity.User `json:"data"`
}
