package main

import (
	"log/slog"
	"os"
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
	repo := repository.NewUserRepository(db)
	userUC := usecase.NewUserUsecase(repo)
	userHandler := handler.NewUserHandler(userUC)

	notificationUC := usecase.NewNotificationUsecase()
	notificationHandler := handler.NewNotificationHandler(notificationUC)

	e := echo.New()
	router.NewRouter(e, *userHandler, *notificationHandler)

	e.Logger.Fatal(e.Start(":8080"))
}