package request

import "found-backend/internal/application/entity"

type CreateUserRequest struct {
	UID      string `json:"uid" validate:"required,min=4,max=20,alphanum"`
	Password string `json:"password" validate:"required,min=8,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required,min=2,max=50"`
}

func (r CreateUserRequest) ToEntity() *entity.User {
	return &entity.User{
		UID:      r.UID,
		Password: r.Password,
		Email:    r.Email,
		Name:     r.Name,
	}
}

type FindUserRequest struct {
	ID int `param:"id"`
}

func (r FindUserRequest) ToEntity() *entity.User {
	return &entity.User{
		ID: r.ID,
	}
}

type UpdateUserRequest struct {
	ID              int    `param:"id"`
	UID             string `json:"uid"`
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

func (r UpdateUserRequest) ToEntity() *entity.User {
	return &entity.User{
		ID:       r.ID,
		UID:      r.UID,
		Password: r.CurrentPassword,
	}
}
