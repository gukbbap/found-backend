package router

import (
	"context"
	"fmt"
	"found-backend/internal/infra/config"

	"go.uber.org/fx"
)

type Router struct {
	config *config.Config
}

func NewRouter(
	lc fx.Lifecycle,
	config *config.Config,

) *Router {
	// init struct
	r := Router{
		config: config,
	}

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					fmt.Println("Hi, it's on start hook of fx")
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return nil
			},
		},
	)

	return &r
}
