package repository

import (
	"context"
	"effective-task/internal/users"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	DELETE_BY_ID = "DELETE FROM users WHERE id = $1"
	ADD          = "INSERT INTO users (name, surname, patronymic, age, gender, nationalize) VALUES (:name, :surname, :patronymic, :age, :gender, :nationalize)"
	UPDATE_BY_ID = `UPDATE users SET name = COALESCE(NULLIF(:name, ''), name), surname = COALESCE(NULLIF(:surname, ''), surname), patronymic = COALESCE(NULLIF(:patronymic, ''), patronymic), age = COALESCE(NULLIF(:age, 0), age), gender = COALESCE(NULLIF(:gender, ''), gender), nationalize = COALESCE(NULLIF(:nationalize, ''), nationalize)
WHERE id = :id
RETURNING *
`
)

type UsersRepo struct {
	db *sqlx.DB
}

func NewUsersRepo(db *sqlx.DB) users.Repository {
	return &UsersRepo{db: db}
}

func (s *UsersRepo) DeleteById(ctx context.Context, id int) error {
	_, err := s.db.ExecContext(ctx, DELETE_BY_ID, id)
	if err != nil {
		return fmt.Errorf("users.repo.DeleteById: %v", err)
	}
	return nil
}

func (s *UsersRepo) UpdateById(ctx context.Context, user *users.User) (*users.User, error) {
	var updatedUser users.User

	req, err := s.db.PrepareNamedContext(ctx, UPDATE_BY_ID)
	if err != nil {
		return nil, fmt.Errorf("users.repo.UpdateById.prepare: %v", err)
	}
	rows, _ := req.QueryxContext(ctx, user)
	defer rows.Close()

	if rows.Next() {
		if err := rows.StructScan(&updatedUser); err != nil {
			return nil, fmt.Errorf("users.repo.UpdateById.scan: %v", err)
		}
	}

	return &updatedUser, nil
}

func (s *UsersRepo) Add(ctx context.Context, user *users.UserDto) error {
	_, err := s.db.NamedExecContext(ctx, ADD, user)
	if err != nil {
		return err
	}
	return nil
}
func (s *UsersRepo) GetAllFiltered(ctx context.Context, filter *users.UserFilter) ([]users.User, error) {

	query, err := filter.CreateQuery()
	if err != nil {
		return nil, fmt.Errorf("users.repo.CreateQuery: %v", err)
	}

	var usersRet []users.User
	if err := s.db.SelectContext(ctx, &usersRet, fmt.Sprintf("SELECT * FROM users %v", query)); err != nil {
		return nil, fmt.Errorf("users.repo.Select: %v", err)
	}

	return usersRet, nil
}
