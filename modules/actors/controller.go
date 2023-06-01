package user

import (
	"crud/dto"
)

type ControllerUser interface {
	CreateUser(req UserParam) (any, error)
	GetUserById(id uint) (FindUser, error)
}

type controllerUser struct {
	userUseCase UseCaseUser
}

func (uc controllerUser) CreateUser(req UserParam) (any, error) {
	user, err := uc.userUseCase.CreateUser(req)
	if err != nil {
		return SuccessCreate{}, err
	}
	res := SuccessCreate{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success create user",
			Message:      "Success Register",
			ResponseTime: "",
		},
		Data: UserParam{
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		},
	}
	return res, nil
}

func (uc controllerUser) GetUserById(id uint) (FindUser, error) {
	var res FindUser
	user, err := uc.userUseCase.GetUserById(id)
	if err != nil {
		return FindUser{}, err
	}
	res.Data = user
	res.ResponseMeta = dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Success create user",
		Message:      "Success Register",
		ResponseTime: "",
	}
	return res, nil
}
