package handler

import "github.com/labstack/echo/v4"

type Handler struct {
	User         UserHandler
	Notification NotificationHandler
}

// NewHandler はHandlerを生成します
func NewHandler(userHandler UserHandler, notificationHandler NotificationHandler) *Handler {
	return &Handler{
		User:         userHandler,
		Notification: notificationHandler,
	}
}

// PostUser は会員登録リクエストをUserHandlerに委譲します
func (h *Handler) PostUser(ctx echo.Context) error {
	return h.User.PostUser(ctx)
}

// PostNotifyReady は焼き上がり通知リクエストをNotificationHandlerに委譲します
func (h *Handler) PostNotifyReady(ctx echo.Context) error {
	return h.Notification.PostNotifyReady(ctx)
}
