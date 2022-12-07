package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type LikesPostgres struct {
	db *sqlx.DB
}

func NewLikesPostgres(db *sqlx.DB) *LikesPostgres {
	return &LikesPostgres{db}
}

func (r *LikesPostgres) IsLiked(postId, userId int) (bool, error) {
	var countOfLikes int
	query := fmt.Sprintf("SELECT count(*) FROM %s WHERE post_id=$1 AND user_id=$2", likesTable)

	err := r.db.Get(&countOfLikes, query, postId, userId)

	return countOfLikes >= 1, err
}

func (r *LikesPostgres) PutLike(postId, userId int) error {
	query := fmt.Sprintf("INSERT INTO %s (post_id, user_id) values ($1, $2)", likesTable)

	_, err := r.db.Exec(query, postId, userId)

	return err
}

func (r *LikesPostgres) Dislike(postId, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE post_id=$1 AND user_id=$2", likesTable)

	_, err := r.db.Exec(query, postId, userId)

	return err
}

func (r *LikesPostgres) GetLikesCountByPostId(postId int) (int, error) {
	var likesCount int
	query := fmt.Sprintf("SELECT count(*) FROM %s WHERE post_id=$1", likesTable)

	err := r.db.Get(&likesCount, query, postId)

	return likesCount, err
}
