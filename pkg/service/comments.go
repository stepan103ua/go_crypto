package service

import (
	"crypto_app/pkg/models"
	"crypto_app/pkg/repository"
)

type CommentsService struct {
	repository repository.Comments
}

func NewCommentsService(repository repository.Comments) *CommentsService {
	return &CommentsService{repository: repository}
}

func (s *CommentsService) GetAllCommentsByPostId(postId int) ([]models.Comment, error) {
	return s.repository.GetAllCommentsByPostId(postId)
}

func (s *CommentsService) CreateComment(comment models.Comment) (int, error) {
	return s.repository.CreateComment(comment)
}

func (s *CommentsService) DeleteComment(commentId, userId int) error {
	return s.repository.DeleteComment(commentId, userId)
}
