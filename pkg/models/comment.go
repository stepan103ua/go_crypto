package models

type Comment struct {
	Id      int    `json:"id" db:"id"`
	PostId  int    `json:"post_id" db:"post_id"`
	Comment string `json:"comment" db:"comment"`
	OwnerId int    `json:"owner_id" db:"owner_id"`
}
