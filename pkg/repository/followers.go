package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type FollowersPostgres struct {
	db *sqlx.DB
}

func NewFollowersPostgres(db *sqlx.DB) *FollowersPostgres {
	return &FollowersPostgres{db}
}

func (r *FollowersPostgres) Follow(followerId, userId int) error {
	query := fmt.Sprintf("INSERT INTO %s (user_id, follower_id) values ($1, $2)", followersTable)

	_, err := r.db.Exec(query, userId, followerId)

	return err
}

func (r *FollowersPostgres) Unfollow(followerId, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1 AND follower_id=$2", followersTable)

	_, err := r.db.Exec(query, userId, followerId)

	return err
}

func (r *FollowersPostgres) IsFollowing(followerId, userId int) (bool, error) {
	var count int
	query := fmt.Sprintf("SELECT count(*) FROM %s WHERE user_id=$1 AND follower_id=$2 LIMIT 1", followersTable)

	err := r.db.Get(&count, query, userId, followerId)

	return count == 1, err
}

func (r *FollowersPostgres) GetFollowersCount(userId int) (int, error) {
	var count int
	query := fmt.Sprintf("SELECT count(*) FROM %s WHERE follower_id=$1", followersTable)

	err := r.db.Get(&count, query, userId)

	return count, err
}
