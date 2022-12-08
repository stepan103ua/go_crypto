package service

import (
	"crypto_app/pkg/repository"
	"errors"
)

type FollowersService struct {
	repository repository.Followers
}

func NewFollowersService(repository repository.Followers) *FollowersService {
	return &FollowersService{repository: repository}
}

func (s *FollowersService) ToggleFollow(followerId, userId int) error {
	if followerId == userId {
		return errors.New("you can not follow yourself")
	}
	isFollowing, err := s.repository.IsFollowing(followerId, userId)

	if err != nil {
		return err
	}

	if isFollowing {
		err = s.repository.Unfollow(followerId, userId)

		if err != nil {
			return err
		}

		return nil
	}

	return s.repository.Follow(followerId, userId)
}

func (s *FollowersService) IsFollowing(followerId, userId int) (bool, error) {
	return s.repository.IsFollowing(followerId, userId)
}

func (s *FollowersService) GetFollowersCount(userId int) (int, error) {
	return s.repository.GetFollowersCount(userId)
}

func (s *FollowersService) GetFollowingCount(userId int) (int, error) {
	return s.repository.GetFollowingCount(userId)
}
