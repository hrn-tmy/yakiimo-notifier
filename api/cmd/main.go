package main

import (
	"log/slog"
	"os"
	"yakiimo-notifier/internal/email"
	"yakiimo-notifier/internal/handler"
	"yakiimo-notifier/internal/infra"
	"yakiimo-notifier/internal/repository"
	"yakiimo-notifier/internal/router"
	"yakiimo-notifier/internal/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := infra.NewDB()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	userRepo := repository.NewUserRepository(db)
	userUC := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUC)

	var mailer email.Sender
	if os.Getenv("EMAIL_DRIVER") == "smtp" {
		mailer = email.NewSMTPSender(
			os.Getenv("SMTP_HOST"),
			os.Getenv("SMTP_PORT"),
			os.Getenv("SMTP_FROM"),
		)
	} else {
		mailer = email.NewSESSender(os.Getenv("SES_FROM"))
	}

	notificationUC := usecase.NewNotificationUsecase(userRepo, mailer)
	notificationHandler := handler.NewNotificationHandler(notificationUC)

	e := echo.New()
	router.NewRouter(e, *userHandler, *notificationHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
