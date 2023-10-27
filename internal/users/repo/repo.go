package repository

import (
	"context"
	"effective-task/internal/users"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UsersRepo struct {
	db *sqlx.DB
}

func NewUsersRepo(db *sqlx.DB) users.Repository {
	return &UsersRepo{db}
}

func (s *UsersRepo) DeleteById(ctx context.Context, id int) error {
	_, err := s.db.ExecContext(ctx, DELETE_BY_ID, id)
	if err != nil {
		return fmt.Errorf("users.repo.DeleteById: %w", err)
	}
	return nil
}

func (s *UsersRepo) UpdateById(ctx context.Context, user *users.User) (users.User, error) {

	req, err := s.db.PrepareNamedContext(ctx, UPDATE_BY_ID)
	if err != nil {
		return users.User{}, fmt.Errorf("users.repo.UpdateById.prepare: %w", err)
	}

	var updatedUser users.User
	if err := req.QueryRowxContext(ctx, user).StructScan(&updatedUser); err != nil {
		return users.User{}, fmt.Errorf("users.repo.UpdateById.scan: %w", err)
	}

	return updatedUser, nil
}

func (s *UsersRepo) Add(ctx context.Context, user *users.User) (users.User, error) {

	query, err := s.db.PrepareNamedContext(ctx, ADD_USER)
	if err != nil {
		return users.User{}, fmt.Errorf("users.repo.add.prepare: %w", err)
	}

	userRet := users.User{}
	if err := query.QueryRowx(user).StructScan(&userRet); err != nil {
		return users.User{}, fmt.Errorf("users.repo.add.scan: %w", err)
	}

	return userRet, nil
}

func (s *UsersRepo) GetAllFiltered(ctx context.Context, filter *users.UserFilter) ([]users.User, error) {

	query, err := filter.CreateQuery()
	if err != nil {
		return nil, fmt.Errorf("users.repo.CreateQuery: %w", err)
	}

	var usersRet []users.User
	if err := s.db.SelectContext(ctx, &usersRet, fmt.Sprintf("SELECT * FROM users %v", query)); err != nil {
		return nil, fmt.Errorf("users.repo.Select: %w", err)
	}

	return usersRet, nil
}
