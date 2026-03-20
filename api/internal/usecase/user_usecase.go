package usecase

import "yakiimo-notifier/internal/repository"

type UserUsecase struct {
	repo repository.IUserRepository
}

func NewUserUsecase(repo repository.IUserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}
