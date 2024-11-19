package service

import (
	"context"
	"found-backend/internal/application/entity"
	"found-backend/internal/application/repository"
	"found-backend/internal/application/service/utils/encrypt"
	"found-backend/internal/infra/exception"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s UserService) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	hashedPassword, err := encrypt.HashPassword(user.Password)
	if err != nil {
		return nil, exception.Wrap(
			err,
			exception.ErrServiceFailedHashPassword,
			exception.StatusInternalServerError,
			"사용자의 비밀번호 암호화에 실패했습니다.",
		)
	}

	user.Password = hashedPassword
	createdUser, err := s.userRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s UserService) FindUser(ctx context.Context, id int) (*entity.User, error) {
	foundUser, err := s.userRepository.FindUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return foundUser, nil
}

func (s UserService) UpdateUser(ctx context.Context, user *entity.User, newPassword string) (*entity.User, error) {

	if newPassword != "" {
		// 기존의 hashed 비밀번호 조회
		foundUser, err := s.userRepository.FindUser(ctx, user.ID)
		if err != nil {
			return nil, err
		}

		err = encrypt.VerifyPassword(user.Password, foundUser.Password)
		if err != nil {
			return nil, exception.Wrap(
				err,
				exception.ErrServiceMismatchPassword,
				exception.StatusBadRequest,
				"비밀번호가 일치하지 않습니다.",
			)
		}

		hashedNewPassword, err := encrypt.HashPassword(newPassword)
		if err != nil {
			return nil, exception.Wrap(
				err,
				exception.ErrServiceFailedHashPassword,
				exception.StatusInternalServerError,
				"새로운 패스워드를 해쉬할 수 없습니다.",
			)
		}

		user.Password = hashedNewPassword
	}

	updatedUser, err := s.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s UserService) DeleteUser(ctx context.Context, id int) error {
	err := s.userRepository.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s UserService) GetLetters(ctx context.Context, userID int) ([]entity.Letter, error) {
	gotLetters, err := s.userRepository.GetLetters(ctx, userID)
	if err != nil {
		return nil, err
	}

	return gotLetters, nil
}

func (s UserService) GetFeelings(ctx context.Context, userID int) ([]entity.Feeling, error) {
	gotFeelings, err := s.userRepository.GetFeelings(ctx, userID)
	if err != nil {
		return nil, err
	}

	return gotFeelings, nil
}
