package handler

import "found-backend/internal/application/service"

type FeelingHandler struct {
	feelingService *service.FeelingService
}

func NewFeelingService(feelingService *service.FeelingService) *FeelingHandler {
	return &FeelingHandler{
		feelingService: feelingService,
	}
}
