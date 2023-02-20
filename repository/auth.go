package repository

import (
	"context"
	"database/sql"
	"payments/models"

	"github.com/zeebo/errs"
)

var authErr = errs.Class("authorization repository error")

type AuthorizationDB struct {
	db *sql.DB
}

func NewAuthorizationDB(db *sql.DB) *AuthorizationDB {
	return &AuthorizationDB{db: db}
}

func (r *AuthorizationDB) CreateUser(ctx context.Context, user models.User) error {
	_, err := r.db.ExecContext(
		ctx, "INSERT INTO users(name, email, password_hash) VALUES ($1, $2, $3);",
		user.Name, user.Email, user.PasswordHash,
	)
	return authErr.Wrap(err)
}

func (r *AuthorizationDB) GetUserID(ctx context.Context, email, passwordHash string) (int, error) {
	var id int
	row := r.db.QueryRowContext(
		ctx, "SELECT id FROM users WHERE email=$1 AND password_hash=$2",
		email, passwordHash,
	)
	if err := row.Scan(&id); err != nil {
		return 0, authErr.Wrap(err)
	}
	return id, nil
}
