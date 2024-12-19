package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/goschool/crud/types"
)

type UserStore interface {
	CreateUser(ctx context.Context, user *types.User) (*types.User, error)
	GetUserByEmail(ctx context.Context, email string) (*types.User, error)
}

type SQLiteUserStore struct {
	db *sql.DB
}

func NewSQLiteUserStore(db *sql.DB) *SQLiteUserStore {
	return &SQLiteUserStore{
		db: db,
	}
}

func (u *SQLiteUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	query := `INSERT INTO users (email, password_hash) VALUES (?, ?)
	RETURNING id`

	var userID string
	err := u.db.QueryRowContext(ctx, query, user.Email, user.PasswordHash).Scan(&userID)

	if err != nil {
		return nil, fmt.Errorf("createUser: %w", err)
	}

	user.ID = userID

	return user, nil
}

func (u *SQLiteUserStore) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	query := `SELECT id, password_hash FROM users WHERE email = ?`

	var user types.User
	err := u.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.PasswordHash)

	if err != nil {
		return nil, fmt.Errorf("getUserByEmail: %w", err)
	}
	user.Email = email

	return &user, nil
}
