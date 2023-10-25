package http

import (
	"effective-task/internal/users"

	"github.com/labstack/echo/v4"
)

func NewUsersRoutes(g echo.Group, h users.Handler) {
	g.GET("", h.GetAllFiltered())
	g.DELETE("/:id", h.DeleteById())
	g.POST("", h.Add())
	g.PATCH("/:id", h.UpdateById())
}
