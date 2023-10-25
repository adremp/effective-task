package http

import (
	"effective-task/internal/users"
	"fmt"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

type usersHandlers struct {
	usersUc users.Usecase
}

func NewUsersHandlers(usersUc users.Usecase) users.Handler {
	return &usersHandlers{usersUc: usersUc}
}

func (s *usersHandlers) DeleteById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Print(err)
			c.JSON(400, "id must be int")
		}

		if err := s.usersUc.DeleteById(c.Request().Context(), id); err != nil {
			log.Print(err)
			c.JSON(500, "error")
		}
		return c.JSON(200, "ok")
	}
}

func (s *usersHandlers) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user users.UserDto
		if err := c.Bind(&user); err != nil {
			log.Print(err)
			c.JSON(400, err)
		}

		if err := s.usersUc.Add(c.Request().Context(), &user); err != nil {
			log.Print(err)
			c.JSON(500, err)
		}

		return c.JSON(200, "ok")
	}
}

func (s *usersHandlers) GetAllFiltered() echo.HandlerFunc {
	return func(c echo.Context) error {
		var filter users.UserFilter
		if err := c.Bind(&filter); err != nil {
			fmt.Print(err)
			c.JSON(400, err)
		}

		users, err := s.usersUc.GetAllFiltered(c.Request().Context(), &filter)
		if err != nil {
			fmt.Print(err)
			c.JSON(500, err)
		}

		return c.JSON(200, users)
	}
}

func (s *usersHandlers) UpdateById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Print(err)
			c.JSON(400, "id must be int")
		}
		var user users.User
		if err := c.Bind(&user); err != nil {
			log.Printf("users.bind: %v", err)
			c.JSON(400, err)
		}
		user.Id = id

		userUpdated, err := s.usersUc.UpdateById(c.Request().Context(), &user)
		if err != nil {
			log.Print(err)
			c.JSON(500, err)
		}

		return c.JSON(200, userUpdated)
	}
}
