package service

import "found-backend/internal/application/repository"

type LetterService struct {
	letterRepository *repository.LetterRepository
}

func NewLetterSerivce(letterRepository *repository.LetterRepository) *LetterService {
	return &LetterService{
		letterRepository: letterRepository,
	}
}
