package userrepo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/iamHemaAI/hema-ai-bot/internal/domain/records"
)

type UserRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *records.User) error {
	_, err := r.db.ExecContext(ctx, queryCreateUser, user.TgID, user.TgUsername, user.Name)
	return err
}

func (r *UserRepository) GetUserByTgID(ctx context.Context, tgID int64) (*records.User, error) {
	row := r.db.QueryRowContext(ctx, queryGetUserByTgID, tgID)

	var user records.User
	err := row.Scan(&user.TgID, &user.TgUsername, &user.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *records.User) error {
	_, err := r.db.ExecContext(ctx, queryUpdateUser, user.TgUsername, user.Name, user.TgID)
	return err
}

func (r *UserRepository) DeleteUser(ctx context.Context, tgID string) error {
	_, err := r.db.ExecContext(ctx, queryDeleteUser, tgID)
	return err
}
