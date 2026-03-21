package usecase

import (
	"context"
	"fmt"
	"os"
	"yakiimo-notifier/internal/repository"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

type NotificationUsecase struct{
	repo repository.IUserRepository
}

func NewNotificationUsecase(repo repository.IUserRepository) *NotificationUsecase {
	return &NotificationUsecase{repo: repo}
}

type SESMailer struct {
	client *ses.Client
	from string
	to []string
}

func (nu *NotificationUsecase) GetTargetUsers(machineID string) ([]string, error) {
	targets, err := nu.repo.GetTargetUsers(machineID)
	if err != nil {
		return nil, err
	}

	return targets, nil
}

func (nu *NotificationUsecase) NotifyReady(to []string, quantity int) error {
	optFns := []func(*config.LoadOptions) error{
		config.WithRegion(os.Getenv("AWS_REGION")),
  }

	// LocalStack用エンドポイントが指定されている場合は上書き
	if endpoint := os.Getenv("AWS_ENDPOINT_URL"); endpoint != "" {
		optFns = append(optFns, config.WithBaseEndpoint(endpoint))
	}
	
	cfg, err := config.LoadDefaultConfig(context.TODO(), optFns...)
	if err != nil {
		return  err
	}

	client := ses.NewFromConfig(cfg)
	subject := "🍠 焼き芋が焼き上がりました！"
	body := fmt.Sprintf("焼き上がり数: %d個\nお早めにどうぞ！",quantity)

	toAddresses := make([]string, len(to))
	copy(toAddresses, to)

	input := &ses.SendEmailInput{
		Source: aws.String(os.Getenv("SES_FROM")),
		Destination: &types.Destination{
			ToAddresses: toAddresses,
		},
		Message: &types.Message{
			Subject: &types.Content{
				Data: aws.String(subject),
			},
			Body: &types.Body{
				Text: &types.Content{
					Data: aws.String(body),
				},
			},
		},
	}

	_, err = client.SendEmail(context.TODO(), input)
	if err != nil {
		return err
	}

	return nil
}