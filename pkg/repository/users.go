package repository

import (
	"crypto_app/pkg/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UsersPostgres struct {
	db *sqlx.DB
}

func NewUsersPostgres(db *sqlx.DB) *UsersPostgres {
	return &UsersPostgres{db}
}

func (r *UsersPostgres) GetUserById(userId int) (models.UserResponse, error) {
	var user models.UserResponse

	query := fmt.Sprintf("SELECT id, name, email, about FROM %s WHERE id=$1 LIMIT 1", usersTable)

	err := r.db.Get(&user, query, userId)

	return user, err
}

func (r *UsersPostgres) UpdateUsername(username string, userId int) error {
	query := fmt.Sprintf("UPDATE %s SET name=$1 WHERE id=$2", usersTable)

	_, err := r.db.Exec(query, username, userId)

	return err
}

func (r *UsersPostgres) UpdateEmail(email string, userId int) error {
	query := fmt.Sprintf("UPDATE %s SET email=$1 WHERE id=$2", usersTable)

	_, err := r.db.Exec(query, email, userId)

	return err
}

func (r *UsersPostgres) UpdatePassword(password string, userId int) error {
	query := fmt.Sprintf("UPDATE %s SET password=$1 WHERE id=$2", usersTable)

	_, err := r.db.Exec(query, password, userId)

	return err
}

func (r *UsersPostgres) UpdateAbout(about string, userId int) error {
	query := fmt.Sprintf("UPDATE %s SET about=$1 WHERE id=$2", usersTable)

	_, err := r.db.Exec(query, about, userId)

	return err
}
