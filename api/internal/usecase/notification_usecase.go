package usecase

import (
	"yakiimo-notifier/internal/repository"
)

type NotificationUsecase struct{
	repo repository.IUserRepository
}

func NewNotificationUsecase(repo repository.IUserRepository) *NotificationUsecase {
	return &NotificationUsecase{repo: repo}
}

func (nu *NotificationUsecase) GetTargetUsers(machineID string) ([]string, error) {
	targets, err := nu.repo.GetTargetUsers(machineID)
	if err != nil {
		return nil, err
	}

	return targets, nil
}