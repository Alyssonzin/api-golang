package usecase

import (
	"api-golang/model"
	"api-golang/repository"
	"fmt"
)

type UserUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return UserUseCase{
		repository: repo,
	}
}

func (uc *UserUseCase) GetUsers() ([]model.User, error) {
	return uc.repository.GetAllUsers()
}

func (uc *UserUseCase) CreateUser(user model.User) (model.User, error) {
	userId, err := uc.repository.CreateUser(user)

	if err != nil {
		fmt.Println(err)
		return model.User{}, err
	}

	user.ID = userId

	return user, nil
}
