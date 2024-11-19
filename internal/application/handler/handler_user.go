package handler

import (
	"found-backend/internal/application/handler/presenter/request"
	"found-backend/internal/application/handler/presenter/response"
	"found-backend/internal/application/handler/utils"
	"found-backend/internal/application/service"
	"net/http"

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
	req, err := utils.ParseRequest[request.CreateUserRequest](c)
	if err != nil {
		return err
	}

	createdUser, err := h.userService.CreateUser(c.Request().Context(), req.ToEntity())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, response.NewCreateUserResponse(createdUser))
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
