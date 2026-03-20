package handler

import "github.com/labstack/echo/v4"

type Handler struct {
	User         UserHandler
	Notification NotificationHandler
}

func NewHandler(userHandler UserHandler, notificationHandler NotificationHandler) *Handler {
	return &Handler{
		User:         userHandler,
		Notification: notificationHandler,
	}
}

func (h *Handler) PostUser(ctx echo.Context) error {
	return h.User.PostUser(ctx)
}

func (h *Handler) GetTargetUsers(ctx echo.Context) error {
	return h.User.GetTargetUsers(ctx)
}

func (h *Handler) PostNotifyReady(ctx echo.Context) error {
	return h.Notification.PostNotifyReady(ctx)
}
