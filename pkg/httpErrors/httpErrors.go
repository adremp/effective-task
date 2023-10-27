package httpErrors

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	ErrNotFound   = errors.New("Not Found")
	ErrValidate   = errors.New("Validation error")
	ErrUnexpected = errors.New("Unexpected error")
	ErrTimeout    = errors.New("Timeout exceeded")
)

type Error struct {
	ErrStatus  int    `json:"status,omitempty"`
	ErrMessage string `json:"message,omitempty"`
}

func ParseError(err error) Error {
	if errors, ok := err.(validator.ValidationErrors); ok {
		mess := fmt.Sprintf("%s %s", strings.ToLower(errors[0].Field()), errors[0].Tag())
		return Error{http.StatusBadRequest, mess}
	}

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return Error{http.StatusNotFound, ErrNotFound.Error()}
	case errors.Is(err, strconv.ErrSyntax):
		return Error{http.StatusBadRequest, ErrValidate.Error()}
	case errors.Is(err, context.DeadlineExceeded):
		return Error{http.StatusRequestTimeout, ErrTimeout.Error()}
	default:
		return Error{http.StatusInternalServerError, ErrUnexpected.Error()}
	}
}

func RequestError(err error) (int, Error) {
	e := ParseError(err)
	return e.ErrStatus, e
}
