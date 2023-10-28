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
	Name        string `db:"name" query:"name"`
	Surname     string `db:"surname" query:"surname"`
	Patronymic  string `db:"patronymic" query:"patronymic"`
	Age         int    `db:"age" query:"age"`
	Gender      string `db:"gender" query:"age"`
	Nationalize string `db:"nationalize" query:"nationalize"`
}

type fF struct {
	idx   int
	field string
}

func formatField(p fF) string {
	return fmt.Sprintf("%v ILIKE $%v", p.field, p.idx)
}

func (f *UserFilter) CreateQuery() (string, []string, error) {
	var queryArr []string
	var valuesArr []string

	idx, format := utils.WithIncreasing(formatField)

	if f.Name != "" {
		queryArr = append(queryArr, format(fF{*idx, "name"}))
		valuesArr = append(valuesArr, f.Name)
	}
	if f.Surname != "" {
		queryArr = append(queryArr, format(fF{*idx, "surname"}))
		valuesArr = append(valuesArr, f.Surname)
	}
	if f.Patronymic != "" {
		queryArr = append(queryArr, format(fF{*idx, "patronymic"}))
		valuesArr = append(valuesArr, f.Patronymic)
	}
	if f.Age != 0 {
		quer, valArr := utils.ParseMinMaxMaybeQuery(*idx, "age", fmt.Sprint(f.Age))
		queryArr = append(queryArr, quer)
		valuesArr = append(valuesArr, valArr...)
	}
	if f.Gender != "" {
		queryArr = append(queryArr, format(fF{*idx, "gender"}))
		valuesArr = append(valuesArr, f.Gender)
	}
	if f.Nationalize != "" {
		queryArr = append(queryArr, format(fF{*idx, "nationalize"}))
		valuesArr = append(valuesArr, f.Nationalize)
	}

	pageQuery := f.PageFilter.CreateQuery()
	if len(queryArr) > 0 {
		queryStr := strings.Join(queryArr, " AND ")
		return fmt.Sprintf("WHERE %v %v", queryStr, pageQuery), valuesArr, nil
	}
	return pageQuery, nil, nil
}
