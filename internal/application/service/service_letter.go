package service

import (
	"context"
	"found-backend/internal/application/entity"
	"found-backend/internal/application/repository"
)

type LetterService struct {
	letterRepository *repository.LetterRepository
}

func NewLetterSerivce(letterRepository *repository.LetterRepository) *LetterService {
	return &LetterService{
		letterRepository: letterRepository,
	}
}

func (s LetterService) CreateLetter(ctx context.Context, letter *entity.Letter) (*entity.Letter, error) {
	createdLetter, err := s.letterRepository.CreateLetter(ctx, letter)
	if err != nil {
		return nil, err
	}

	return createdLetter, nil
}

func (s LetterService) FindLetter(ctx context.Context, id int) (*entity.Letter, error) {
	foundLetter, err := s.letterRepository.FindLetter(ctx, id)
	if err != nil {
		return nil, err
	}

	return foundLetter, nil
}

func (s LetterService) UpdateLetter(ctx context.Context, letter *entity.Letter) (*entity.Letter, error) {
	updatedLetter, err := s.letterRepository.UpdateLetter(ctx, letter)
	if err != nil {
		return nil, err
	}

	return updatedLetter, err
}

func (s LetterService) DeleteLetter(ctx context.Context, id int) error {
	err := s.letterRepository.DeleteLetter(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s LetterService) GetFeelings(ctx context.Context, userID int) ([]entity.Feeling, error) {
	gotFeelings, err := s.letterRepository.GetFeelings(ctx, userID)
	if err != nil {
		return nil, err
	}

	return gotFeelings, nil
}
