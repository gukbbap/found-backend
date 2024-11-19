package router

import (
	"context"
	"found-backend/internal/application/handler"
	"found-backend/internal/infra/config"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type Router struct {
	config *config.Config

	// handlers
	userHandler *handler.UserHandler

	// router
	engine *echo.Echo
}

func NewRouter(
	lc fx.Lifecycle,

	config *config.Config,
	userHandler *handler.UserHandler,
) *Router {
	// init struct
	r := Router{
		config: config,

		// handlers
		userHandler: userHandler,

		// engine
		engine: echo.New(),
	}

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					if err := r.Listen(r.config.Port); err != nil {
						panic(err)
					}
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

func (r *Router) route() {
	api := r.engine.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			users := v1.Group("/users")
			{
				r.userHandler.Route(users)
			}
		}
	}
}

func (r *Router) Listen(addr string) error {
	return r.engine.Start(addr)
}
