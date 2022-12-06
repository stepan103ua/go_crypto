package service

import (
	"crypto_app/pkg/models"
	"crypto_app/pkg/repository"
)

type RepliesService struct {
	repository repository.Replies
}

func NewRepliesService(repository repository.Replies) *RepliesService {
	return &RepliesService{repository: repository}
}

func (s *RepliesService) CreateReply(reply models.Reply) (int, error) {
	return s.repository.CreateReply(reply)
}

func (s *RepliesService) GetRepliesByCommentId(commentId int) ([]models.Reply, error) {
	return s.repository.GetRepliesByCommentId(commentId)
}
