package middleware

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

func RequestLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)

		req := c.Request()
		res := c.Response()

		diff := time.Since(start)
		log.Printf("%s | %v | %s | %s", req.Method, res.Status, req.URL, diff)
		return err
	}
}
