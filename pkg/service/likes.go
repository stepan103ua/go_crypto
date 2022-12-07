package service

import "crypto_app/pkg/repository"

type LikesService struct {
	repository repository.Likes
}

func NewLikesService(repository repository.Likes) *LikesService {
	return &LikesService{repository: repository}
}

func (s *LikesService) ToggleLike(postId, userId int) error {
	isLiked, err := s.repository.IsLiked(postId, userId)
	if err != nil {
		return err
	}

	if isLiked {
		err = s.repository.Dislike(postId, userId)
		if err != nil {
			return err
		}
		return nil
	}

	err = s.repository.PutLike(postId, userId)

	return err
}

func (s *LikesService) GetLikesCountByPostId(postId int) (int, error) {
	return s.repository.GetLikesCountByPostId(postId)
}
