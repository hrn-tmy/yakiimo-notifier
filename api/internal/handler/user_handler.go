package handler

import (
	"net/http"
	"yakiimo-notifier/internal/constant"
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
		return ErrorResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	message := validateCreateUser(req)
	if message != "" {
		return ErrorResponse(ctx, message, http.StatusBadRequest)
	}

	user, err := uh.uc.CreateUser(string(req.Email), req.Name, req.Password)
	if err != nil {
		return ErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
	}

	res := gen.CreateUserResponse{
		Status: constant.SuccessStatus,
		User:   user,
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (uh *UserHandler) GetTargetUsers(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}
