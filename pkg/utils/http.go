package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)


func SanitizeRequest[T any](c echo.Context, request *T) error {
	if err := c.Bind(request); err != nil {
		return err
	}

	if err := Validate.Struct(request); err != nil {
		return err.(validator.ValidationErrors)
	}

	return nil
}