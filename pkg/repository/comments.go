package repository

import (
	"crypto_app/pkg/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type CommentsPostgres struct {
	db *sqlx.DB
}

func NewCommentsPostgres(db *sqlx.DB) *CommentsPostgres {
	return &CommentsPostgres{db}
}

func (c *CommentsPostgres) GetAllCommentsByPostId(postId int) ([]models.Comment, error) {
	var comments []models.Comment = make([]models.Comment, 0)

	query := fmt.Sprintf("SELECT * FROM %s WHERE post_id=$1", commentsTable)

	err := c.db.Select(&comments, query, postId)

	return comments, err
}
