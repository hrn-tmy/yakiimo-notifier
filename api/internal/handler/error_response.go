package handler

import (
	"yakiimo-notifier/internal/constant"
	"yakiimo-notifier/internal/gen"

	"github.com/labstack/echo/v4"
)

// ErrorResponse はエラーレスポンスをJSON形式で返す
func ErrorResponse(ctx echo.Context, message string, httpStatus int) error {
	res := gen.ErrorResponse{
		Status:  constant.ErrorStatus,
		Message: message,
	}
	return ctx.JSON(httpStatus, res)
}
