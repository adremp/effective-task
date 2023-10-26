package utils

import "github.com/labstack/echo/v4"


func SanitizeRequest[T any](c echo.Context, request *T) error {
	if err := c.Bind(request); err != nil {
		return err
	}

	if err := Validate.Struct(request); err != nil {
		return err
	}

	return nil
}