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

	query := fmt.Sprintf("SELECT comments.*, users.name FROM %s JOIN users ON comments.owner_id=users.id WHERE post_id=$1", commentsTable)

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

func (c *CommentsPostgres) DeleteComment(commentId, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND owner_id=$2", commentsTable)

	_, err := c.db.Exec(query, commentId, userId)

	return err
}

func (c *CommentsPostgres) UpdateComment(comment string, commentId, userId int) error {
	query := fmt.Sprintf("UPDATE %s SET comment=$1 WHERE id=$2 AND owner_id=$3", commentsTable)

	_, err := c.db.Exec(query, comment, commentId, userId)

	return err
}

func (c *CommentsPostgres) GetCommentsCountByPostId(postId int) (int, error) {
	var commentsCount int = 0
	queryComments := "select count(*) + (select count(*) from comments join replies on replies.comment_id=comments.id where comments.post_id=$1) from comments where post_id=$1"

	err := c.db.Get(&commentsCount, queryComments, postId)

	return commentsCount, err
}
