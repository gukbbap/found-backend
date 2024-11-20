package response

import (
	"found-backend/internal/application/entity"
	"time"
)

type CreateUserResponse struct {
	ID        int       `json:"id"`
	UID       string    `json:"uid"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewCreateUserResponse(user *entity.User) *CreateUserResponse {
	return &CreateUserResponse{
		ID:        user.ID,
		UID:       user.UID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}
}

type FindUserResponse struct {
	ID        int       `json:"id"`
	UID       string    `json:"uid"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewFindUserResponse(user *entity.User) *FindUserResponse {
	return &FindUserResponse{
		ID:        user.ID,
		UID:       user.UID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type UpdateUserResponse struct {
	ID        int       `json:"id"`
	UID       string    `json:"uid"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewUpdateUserResponse(user *entity.User) *UpdateUserResponse {
	return &UpdateUserResponse{
		ID:        user.ID,
		UID:       user.UID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type DeleteUserResponse struct {
	ID    int    `json:"id"`
	UID   string `json:"uid"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func NewDeleteUserResponse(user *entity.User) *DeleteUserResponse {
	return &DeleteUserResponse{
		ID:    user.ID,
		UID:   user.UID,
		Email: user.Email,
		Name:  user.Name,
	}
}

type GetLettersResponse struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

func NewGetLettersResponse(letters []entity.Letter) []GetLettersResponse {
	var resp []GetLettersResponse
	for _, letter := range letters {
		resp = append(resp, GetLettersResponse{
			ID:      letter.ID,
			Content: letter.Content,
		})
	}

	return resp
}
