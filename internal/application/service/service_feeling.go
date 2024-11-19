package service

import (
	"context"
	"found-backend/internal/application/entity"
	"found-backend/internal/application/repository"
)

type FeelingService struct {
	feelingRepository *repository.FeelingRepository
}

func NewFeelingService(feelingRepository *repository.FeelingRepository) *FeelingService {
	return &FeelingService{
		feelingRepository: feelingRepository,
	}
}

func (s FeelingService) CreateFeeling(ctx context.Context, feeling *entity.Feeling) (*entity.Feeling, error) {
	createdFeeling, err := s.feelingRepository.CreateFeeling(ctx, feeling)
	if err != nil {
		return nil, err
	}

	return createdFeeling, nil
}

func (s FeelingService) FindFeeling(ctx context.Context, id int) (*entity.Feeling, error) {
	foundFeeling, err := s.feelingRepository.FindFeeling(ctx, id)
	if err != nil {
		return nil, err
	}

	return foundFeeling, nil
}

func (s FeelingService) UpdateFeeling(ctx context.Context, feeling *entity.Feeling) (*entity.Feeling, error) {
	updatedFeeling, err := s.feelingRepository.UpdateFeeling(ctx, feeling)
	if err != nil {
		return nil, err
	}

	return updatedFeeling, nil
}

func (s FeelingService) DeleteFeeling(ctx context.Context, id int) error {
	err := s.feelingRepository.DeleteFeeling(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
