package handler

import (
	"found-backend/internal/application/service"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h UserHandler) Route(g *echo.Group) {
	g.POST("/", h.CreateUser)
	g.GET("/:id", h.FindUser)
	g.PATCH("/:id", h.UpdateUser)
	g.DELETE("/:id", h.DeleteUser)

	g.GET("/:id/letters", h.GetLetters)
	g.GET("/:id/feelings", h.GetFeelings)
}

func (h UserHandler) CreateUser(c echo.Context) error {
	return nil
}

func (h UserHandler) FindUser(c echo.Context) error {
	return nil
}

func (h UserHandler) UpdateUser(e echo.Context) error {
	return nil
}

func (h UserHandler) DeleteUser(e echo.Context) error {
	return nil
}

func (h UserHandler) GetLetters(e echo.Context) error {
	return nil
}

func (h UserHandler) GetFeelings(e echo.Context) error {
	return nil
}
