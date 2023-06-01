package user

import (
	"crud/entity"
	"crud/repository"
	"time"
)

type UseCaseUser interface {
	CreateUser(user UserParam) (entity.User, error)
	GetUserById(id uint) (entity.User, error)
}

type useCaseUser struct {
	userRepo repository.UserInterfaceRepo
}

func (uc useCaseUser) CreateUser(user UserParam) (entity.User, error) {
	var newUser *entity.User

	newUser = &entity.User{
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := uc.userRepo.CreateUser(newUser)
	if err != nil {
		return *newUser, err
	}
	return *newUser, nil
}

func (uc useCaseUser) GetUserById(id uint) (entity.User, error) {
	var user entity.User
	user, err := uc.userRepo.GetUserById(id)
	return user, err
}
