package psql

import (
	"context"
	"database/sql"

	"github.com/awleory/kode/notebook/internal/entity"
)

type Users struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) *Users {
	return &Users{
		db: db,
	}
}

func (r *Users) CreateUser(ctx context.Context, user entity.SignUpInput) error {
	_, err := r.db.Exec("INSERT INTO users (email, password_hash) values ($1, $2)",
		user.Email, user.Password)

	return err
}

func (r *Users) GetUser(ctx context.Context, email, password string) (int, error) {
	var userId int
	err := r.db.QueryRow("SELECT id FROM users WHERE email=$1 AND password_hash=$2", email, password).
		Scan(&userId)

	return userId, err
}
