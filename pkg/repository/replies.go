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

func (r *RepliesPostgres) GetRepliesByCommentId(commentId int) ([]models.Reply, error) {
	var replies []models.Reply = make([]models.Reply, 0)

	query := fmt.Sprintf("SELECT * FROM %s WHERE comment_id=$1", repliesTable)

	err := r.db.Select(&replies, query, commentId)

	return replies, err
}
