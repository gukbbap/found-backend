package response

import (
	"found-backend/internal/application/entity"
	"time"
)

type CreateUserResponse struct {
	UID       string    `json:"uid"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewCreateUserResponse(user *entity.User) *CreateUserResponse {
	return &CreateUserResponse{
		UID:       user.UID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}
}

type FindUserResponse struct {
	UID       string    `json:"uid"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewFindUserResponse(user *entity.User) *FindUserResponse {
	return &FindUserResponse{
		UID:       user.UID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type UpdateUserResponse struct {
	UID       string    `json:"uid"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewUpdateUserResponse(user *entity.User) *UpdateUserResponse {
	return &UpdateUserResponse{
		UID:       user.UID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type DeleteUserResponse struct {
	UID   string `json:"uid"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func NewDeleteUserResponse(user *entity.User) *DeleteUserResponse {
	return &DeleteUserResponse{
		UID:   user.UID,
		Email: user.Email,
		Name:  user.Name,
	}
}
