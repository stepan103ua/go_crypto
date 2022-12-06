package models

type Reply struct {
	Id        int    `json:"id" db:"id"`
	CommentId int    `json:"comment_id" db:"comment_id"`
	OwnerId   int    `json:"owner_id" db:"owner_id"`
	Reply     string `json:"reply" db:"reply" binding:"required"`
}
