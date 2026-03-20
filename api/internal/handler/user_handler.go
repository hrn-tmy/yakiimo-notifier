package handler

import (
	"net/http"
	"yakiimo-notifier/internal/usecase"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	uc *usecase.UserUsecase
}

func NewUserHandler(uc *usecase.UserUsecase) *UserHandler {
	return &UserHandler{uc: uc}
}

func (uh *UserHandler) GetTargetUsers(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}
