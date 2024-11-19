package main

import (
	router "found-backend/internal"
	"found-backend/internal/application"
	"found-backend/internal/infra"
	"found-backend/internal/infra/storage/mysql"
	"log"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Supply(".env.toml"),
		fx.Provide(
			router.NewRouter,
		),

		// infra modules
		infra.ConfigModule,
		infra.StorageModule,

		// application modules
		application.RepositoryModule,
		application.ServiceModule,
		application.HandlerModule,

		fx.Invoke(
			func(m *mysql.MySQL) {
				if err := m.AutoMigration(); err != nil {
					log.Panicf("failed to auto migration: %v\n", err)
				}
			},
			func(r *router.Router) {},
		),
	).Run()
}
