package application

import (
	"found-backend/internal/application/handler"
	"found-backend/internal/application/repository"
	"found-backend/internal/application/service"

	"go.uber.org/fx"
)

var RepositoryModule = fx.Module(
	"repository-module",
	fx.Provide(
		repository.NewUserRepository,
		repository.NewLetterRepository,
		repository.NewFeelingRepository,
	),
)

var ServiceModule = fx.Module(
	"service-module",
	fx.Provide(
		service.NewUserService,
		service.NewLetterSerivce,
		service.NewFeelingService,
	),
)

var HandlerModule = fx.Module(
	"handler-module",
	fx.Provide(
		handler.NewUserHandler,
		handler.NewLetterHandler,
		handler.NewFeelingService,
	),
)
