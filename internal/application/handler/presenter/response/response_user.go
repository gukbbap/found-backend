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
