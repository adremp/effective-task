package users

import (
	"context"
	"effective-task/pkg/utils"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	DeleteById() echo.HandlerFunc
	UpdateById() echo.HandlerFunc
	Add() echo.HandlerFunc
	GetAllFiltered() echo.HandlerFunc
}

type Usecase interface {
	DeleteById(context.Context, int) error
	UpdateById(context.Context, *User) (User, error)
	Add(context.Context, *User) (User, error)
	GetAllFiltered(context.Context, *UserFilter) ([]User, error)
}

type Repository interface {
	DeleteById(context.Context, int) error
	UpdateById(context.Context, *User) (User, error)
	Add(context.Context, *User) (User, error)
	GetAllFiltered(context.Context, *UserFilter) ([]User, error)
}

type User struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Surname     string `db:"surname"`
	Patronymic  string `db:"patronymic"`
	Age         int    `db:"age"`
	Gender      string `db:"gender"`
	Nationalize string `db:"nationalize"`
}

type UserFilter struct {
	utils.PageFilter
	Name        string `db:"name"`
	Surname     string `db:"surname"`
	Patronymic  string `db:"patronymic"`
	Age         int    `db:"age"`
	Gender      string `db:"gender"`
	Nationalize string `db:"nationalize"`
}

func (f *UserFilter) CreateQuery() (string, error) {
	var queryArr []string

	if f.Name != "" {
		queryArr = append(queryArr, fmt.Sprintf("name ILIKE '%v'", f.Name))
	}
	if f.Surname != "" {
		queryArr = append(queryArr, fmt.Sprintf("surname ILIKE '%v'", f.Surname))
	}
	if f.Patronymic != "" {
		queryArr = append(queryArr, fmt.Sprintf("patronymic ILIKE '%v'", f.Patronymic))
	}
	if f.Age != 0 {
		queryArr = append(queryArr, utils.ParseMinMaxMaybeQuery("age", fmt.Sprintf("%v", f.Age)))
	}
	if f.Gender != "" {
		queryArr = append(queryArr, fmt.Sprintf("gender ILIKE '%v'", f.Gender))
	}
	if f.Nationalize != "" {
		queryArr = append(queryArr, fmt.Sprintf("nationalize ILIKE '%v'", f.Nationalize))
	}

	queryStr := strings.Join(queryArr, " AND ")
	pageQuery := f.PageFilter.WithQuery(queryStr)
	if queryStr != "" {
		return fmt.Sprintf("WHERE %v", pageQuery), nil
	}
	return pageQuery, nil
}
