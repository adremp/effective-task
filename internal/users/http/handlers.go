package http

import (
	"effective-task/internal/users"
	"effective-task/pkg/utils"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

type usersHandlers struct {
	usersUc users.Usecase
}

func NewUsersHandlers(usersUc users.Usecase) users.Handler {
	return &usersHandlers{usersUc}
}

func (s *usersHandlers) DeleteById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Print(err)
			return c.JSON(400, "id must be int")
		}

		if err := s.usersUc.DeleteById(c.Request().Context(), id); err != nil {
			log.Print(err)
			return c.JSON(500, "error")
		}
		return c.JSON(200, "ok")
	}
}

func (s *usersHandlers) Add() echo.HandlerFunc {
	type UserAdd struct {
		users.User
		Name    string `validate:"required"`
		Surname string `validate:"required"`
	}
	return func(c echo.Context) error {
		var user UserAdd
		if err := utils.SanitizeRequest(c, &user); err != nil {
			log.Print(err)
			return c.JSON(400, err)
		}

		userToAdd := user.User
		userToAdd.Name = user.Name
		userToAdd.Surname = user.Surname

		userRet, err := s.usersUc.Add(c.Request().Context(), &userToAdd)
		if err != nil {
			log.Print(err)
			return c.JSON(500, err)
		}

		return c.JSON(200, userRet)
	}
}

func (s *usersHandlers) GetAllFiltered() echo.HandlerFunc {
	return func(c echo.Context) error {
		var filter users.UserFilter
		if err := utils.SanitizeRequest(c, &filter); err != nil {
			log.Print(err)
			return c.JSON(400, err)
		}

		users, err := s.usersUc.GetAllFiltered(c.Request().Context(), &filter)
		if err != nil {
			log.Print(err)
			return c.JSON(500, err)
		}

		return c.JSON(200, users)
	}
}

func (s *usersHandlers) UpdateById() echo.HandlerFunc {
	type UpdateUser struct {
		users.User
		Id int `validate:"required"`
	}
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		user := UpdateUser{Id: id}
		if err := utils.SanitizeRequest(c, &user); err != nil {
			log.Print(err)
			return c.JSON(400, err)
		}

		userToUpdate := user.User
		userToUpdate.Id = user.Id

		userUpdated, err := s.usersUc.UpdateById(c.Request().Context(), &userToUpdate)
		if err != nil {
			log.Print(err)
			return c.JSON(500, err)
		}

		return c.JSON(200, userUpdated)
	}
}
