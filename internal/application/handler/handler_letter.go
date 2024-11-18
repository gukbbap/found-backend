package handler

import "found-backend/internal/application/service"

type LetterHandler struct {
	letterService *service.LetterService
}

func NewLetterHandler(letterService *service.LetterService) *LetterHandler {
	return &LetterHandler{
		letterService: letterService,
	}
}
