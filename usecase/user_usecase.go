package usecase

import (
	"api-golang/model"
	"api-golang/repository"
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
