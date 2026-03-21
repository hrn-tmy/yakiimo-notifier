package router

import (
	"yakiimo-notifier/internal/gen"
	"yakiimo-notifier/internal/handler"

	"github.com/labstack/echo/v4"
)

// NewRouter　はルーティングの初期化を行います
func NewRouter(e *echo.Echo, userHandler handler.UserHandler, notificationHandler handler.NotificationHandler) {
	h := handler.NewHandler(userHandler, notificationHandler)
	wrapper := gen.ServerInterfaceWrapper{
		Handler: h,
	}

	user := e.Group("/users")
	user.POST("", wrapper.PostUser)

	notify := e.Group("/notify")
	notify.POST("/ready", wrapper.PostNotifyReady)
}
