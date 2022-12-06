package repository

import (
	"crypto_app/pkg/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type RepliesPostgres struct {
	db *sqlx.DB
}

func NewRepliesPostgres(db *sqlx.DB) *RepliesPostgres {
	return &RepliesPostgres{db}
}

func (r *RepliesPostgres) CreateReply(reply models.Reply) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (comment_id, owner_id, reply) values ($1, $2, $3) RETURNING id", repliesTable)

	row := r.db.QueryRow(query, reply.CommentId, reply.OwnerId, reply.Reply)

	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
