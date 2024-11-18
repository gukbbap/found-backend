package service

import "found-backend/internal/application/repository"

type FeelingService struct {
	feelingRepository *repository.FeelingRepository
}

func NewFeelingService(feelingRepository *repository.FeelingRepository) *FeelingService {
	return &FeelingService{
		feelingRepository: feelingRepository,
	}
}
