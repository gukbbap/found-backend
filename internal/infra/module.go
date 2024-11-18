package infra

import (
	"found-backend/internal/infra/config"
	"found-backend/internal/infra/storage/mysql"

	"go.uber.org/fx"
)

var ConfigModule = fx.Module(
	"config-moudle",
	fx.Provide(
		config.NewConfig,
	),
)

var StorageModule = fx.Module(
	"storage-module",
	fx.Provide(
		mysql.NewMySQL,
	),
)
