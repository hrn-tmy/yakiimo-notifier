package handler

import (
	"net/http"
	"yakiimo-notifier/internal/gen"
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
	var req gen.PostNotifyReadyJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		return ErrorResponse(ctx, err.Error(), http.StatusBadRequest)
	}

	message := validateNotifyReady(req)
	if message != "" {
		return ErrorResponse(ctx, message, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK, nil)
}
