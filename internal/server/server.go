package server

import (
	"effective-task/internal/middleware"
	"effective-task/internal/users/http"
	repository "effective-task/internal/users/repo"
	"effective-task/internal/users/usecase"
	"effective-task/pkg/postgres"
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

func NewServer() error {
	e := echo.New()
	e.Server.ReadTimeout = 3 * time.Second
	e.Server.WriteTimeout = 3 * time.Second

	e.Use(middleware.RequestLogger)
	log.SetFlags(log.Ldate)

	psqlDB, err := postgres.NewPgConn()
	if err != nil {
		return fmt.Errorf("failed to connect to postgres: %w", err)
	}
	defer psqlDB.Close()

	v1 := e.Group("/api/v1")

	usersRepo := repository.NewUsersRepo(psqlDB)
	usersUC := usecase.NewUsersUc(usersRepo)
	usersH := http.NewUsersHandlers(usersUC)

	http.NewUsersRoutes(*v1.Group("/users"), usersH)

	e.Logger.Fatal(e.Start(":8080"))
	return nil
}
