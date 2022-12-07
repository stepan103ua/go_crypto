package repository

import (
	"crypto_app/pkg/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PostsPostgres struct {
	db *sqlx.DB
}

func NewPostsPostgres(db *sqlx.DB) *PostsPostgres {
	return &PostsPostgres{db}
}

func (p *PostsPostgres) CreatePost(post models.Post) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, description, image_url, crypto_currency, created_at, owner_id) values ($1, $2, $3, $4, NOW(), $5) RETURNING id", postsTable)

	row := p.db.QueryRow(query, post.Title, post.Description, post.ImageUrl, post.CryptoCurrency, post.UserId)

	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (p *PostsPostgres) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	query := fmt.Sprintf("SELECT posts.*, users.name FROM %s JOIN users ON posts.owner_id=users.id WHERE posts.owner_id=$1", postsTable)

	err := p.db.Select(&posts, query)

	return posts, err
}

func (p *PostsPostgres) GetAllPostsByUserId(userId int) ([]models.Post, error) {
	var posts []models.Post
	query := fmt.Sprintf("SELECT posts.*, users.name FROM %s JOIN users ON posts.owner_id=users.id WHERE posts.owner_id=$1", postsTable)

	err := p.db.Select(&posts, query, userId)

	return posts, err
}

func (p *PostsPostgres) DeletePost(postId, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND owner_id=$2", postsTable)

	_, err := p.db.Exec(query, postId, userId)

	return err
}

func (p *PostsPostgres) UpdatePost(title, description, cryptoCurrency, image_url string, postId, userId int) error {
	query := fmt.Sprintf("UPDATE %s SET title=$1, description=$2, crypto_currency=$3, image_url=$4 WHERE id=$5 AND owner_id=$6", postsTable)

	_, err := p.db.Exec(query, title, description, cryptoCurrency, image_url, postId, userId)

	return err
}
