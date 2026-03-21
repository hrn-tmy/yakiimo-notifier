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

// NewNotificationUsecase はNotificationUsecaseを生成します
func NewNotificationUsecase(repo repository.IUserRepository, mailer email.Sender) *NotificationUsecase {
	return &NotificationUsecase{repo: repo, mailer: mailer}
}

// GetTargetUsers は指定した機械IDに対して通知対象となる会員のメールアドレス一覧を返す
func (nu *NotificationUsecase) GetTargetUsers(machineID string) ([]string, error) {
	targets, err := nu.repo.GetTargetUsers(machineID)
	if err != nil {
		return nil, err
	}

	return targets, nil
}

// NotifyReady は焼き上がり通知メールを対象会員全員に送信します
func (nu *NotificationUsecase) NotifyReady(to []string, quantity int) error {
	subject := "🍠 焼き芋が焼き上がりました！"
	body := fmt.Sprintf("焼き上がり数: %d個\nお早めにどうぞ！", quantity)
	return nu.mailer.Send(to, subject, body)
}
