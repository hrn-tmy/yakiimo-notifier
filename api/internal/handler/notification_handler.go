package handler

import (
	"net/http"
	"yakiimo-notifier/internal/usecase"

	"github.com/labstack/echo/v4"
)

type NotificationHandler struct {
	uc *usecase.NotificationUsecase
}

func NewNotificationHandler(uc *usecase.NotificationUsecase) *NotificationHandler {
	return &NotificationHandler{uc: uc}
}

func (nh *NotificationHandler) PostNotifyReady(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}