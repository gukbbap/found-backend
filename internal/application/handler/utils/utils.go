package utils

import (
	"found-backend/internal/infra/exception"

	"github.com/labstack/echo/v4"
)

func ParseRequest[T any](c echo.Context) (*T, error) {
	req := new(T)

	if err := c.Bind(req); err != nil {
		return nil, exception.Wrap(
			err,
			exception.ErrPresenterFailedBind,
			exception.StatusInternalServerError,
			"해당 요청을 bind하는데 실패했습니다.",
		)
	}

	if err := c.Validate(req); err != nil {
		return nil, exception.Wrap(
			err,
			exception.ErrPresenterFailedValidation,
			exception.StatusBadRequest,
			"해당 요청을 validate하는데 실패했습니다.",
		)
	}

	return req, nil
}
