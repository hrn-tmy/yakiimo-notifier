package handler

import (
	"yakiimo-notifier/internal/constant"
	"yakiimo-notifier/internal/gen"

	"github.com/labstack/echo/v4"
)

func ErrorResponse(ctx echo.Context, httpStatus int) error {
	res := gen.ErrorResponse{
		Status: constant.ErrorStatus,
		Message: "処理に失敗しました。",
	}
	return ctx.JSON(httpStatus, res)
}