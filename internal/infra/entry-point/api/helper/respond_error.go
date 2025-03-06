package helper

import (
	"errors"
	"net/http"

	"github.com/carlosblanco00/order-manager-api/internal/domain/model"
	"github.com/labstack/echo/v4"
)

const InternalServerErrorMessage = "Internal server error"

func RespondError(c echo.Context, err error) error {
	var appErr model.AppError
	if errors.As(err, &appErr) {
		if status, ok := ErrCodeMapping[appErr.Code]; ok {

			return c.JSON(status, appErr)
		}
	}
	return c.JSON(http.StatusInternalServerError, model.AppError{Code: model.ErrCodeInternalServerError, Msg: InternalServerErrorMessage})

}

var ErrCodeMapping map[string]int = map[string]int{
	model.ErrCodeIdenpotency:       http.StatusConflict,
	model.ErrCodeNotFound:          http.StatusNotFound,
	model.ErrCodeInvalidParams:     http.StatusBadRequest,
	model.ErrCodeBadRequest:        http.StatusBadRequest,
	model.ErrCodeHeaderIdenpotency: http.StatusBadRequest,
	model.ErrCodeStock:             http.StatusConflict,
}
