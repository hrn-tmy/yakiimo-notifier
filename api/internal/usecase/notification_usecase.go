package usecase

import (
	"fmt"
	"yakiimo-notifier/internal/email"
	"yakiimo-notifier/internal/repository"
)

type NotificationUsecase struct {
	repo   repository.IUserRepository
	mailer email.Sender
}

func NewNotificationUsecase(repo repository.IUserRepository, mailer email.Sender) *NotificationUsecase {
	return &NotificationUsecase{repo: repo, mailer: mailer}
}

func (nu *NotificationUsecase) GetTargetUsers(machineID string) ([]string, error) {
	targets, err := nu.repo.GetTargetUsers(machineID)
	if err != nil {
		return nil, err
	}

	return targets, nil
}

func (nu *NotificationUsecase) NotifyReady(to []string, quantity int) error {
	subject := "🍠 焼き芋が焼き上がりました！"
	body := fmt.Sprintf("焼き上がり数: %d個\nお早めにどうぞ！", quantity)
	return nu.mailer.Send(to, subject, body)
}
