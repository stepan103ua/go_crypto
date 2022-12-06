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

func (c *CommentsPostgres) CreateComment(comment models.Comment) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (post_id, comment, owner_id) values ($1, $2, $3) RETURNING id", commentsTable)

	row := c.db.QueryRow(query, comment.PostId, comment.Comment, comment.OwnerId)

	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, err
}
