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

	query := fmt.Sprintf("SELECT id, name, email FROM %s WHERE id=$1 LIMIT 1", usersTable)

	err := r.db.Get(&user, query, userId)

	return user, err
}
