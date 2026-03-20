package handler

import (
	"net/http"
	"yakiimo-notifier/internal/gen"
	"yakiimo-notifier/internal/usecase"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	uc *usecase.UserUsecase
}

func NewUserHandler(uc *usecase.UserUsecase) *UserHandler {
	return &UserHandler{uc: uc}
}

func (uh *UserHandler) PostUser(ctx echo.Context) error {
	var req gen.CreateUserRequest
	if err := ctx.Bind(&req); err != nil {
		return ErrorResponse(ctx, http.StatusBadRequest)
	}
	message := validateCreateUser(req)
	if message != "" {
		return ErrorResponse(ctx, http.StatusBadRequest)
	}
	return ctx.JSON(http.StatusCreated, nil)
}

func (uh *UserHandler) GetTargetUsers(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}
