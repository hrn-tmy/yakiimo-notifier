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

// NewNotificationHandler はNotificationHandlerを生成します
func NewNotificationHandler(uc *usecase.NotificationUsecase) *NotificationHandler {
	return &NotificationHandler{uc: uc}
}

// PostNotifyReady は焼き上がり通知リクエストを処理し、対象会員にメールを送信します
func (nh *NotificationHandler) PostNotifyReady(ctx echo.Context) error {
	var req gen.PostNotifyReadyJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		return ErrorResponse(ctx, err.Error(), http.StatusBadRequest)
	}

	message := validateNotifyReady(req)
	if message != "" {
		return ErrorResponse(ctx, message, http.StatusBadRequest)
	}

	to, err := nh.uc.GetTargetUsers(req.MachineId)
	if err != nil {
		return ErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
	}

	if err := nh.uc.NotifyReady(to, 10); err != nil {
		return ErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK, nil)
}
