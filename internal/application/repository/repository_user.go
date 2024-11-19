package repository

import (
	"context"
	"found-backend/internal/application/entity"
	"found-backend/internal/infra/exception"
	"found-backend/internal/infra/storage/mysql"
	"found-backend/pkg/defaults"
)

type UserRepository struct {
	db *mysql.MySQL
}

func NewUserRepository(db *mysql.MySQL) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r UserRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	err := r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return nil, exception.Wrap(
			err,
			exception.ErrMySQLInternal,
			exception.StatusInternalServerError,
			"사용자를 생성할 수 없습니다.",
		)
	}

	return user, nil
}

func (r UserRepository) FindUser(ctx context.Context, id int) (*entity.User, error) {
	var foundUser entity.User
	err := r.db.WithContext(ctx).First(&foundUser).Error
	if err != nil {
		return nil, exception.Wrap(
			err,
			exception.ErrMySQLNotFound,
			exception.StatusNotFound,
			"사용자를 조회할 수 없습니다.",
		)
	}

	return &foundUser, nil
}

func (r UserRepository) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	// Partial Update
	updateFields := map[string]any{}

	// 업데이트를 위한 필드 검사
	if user.UID != defaults.String {
		updateFields["uid"] = user.UID
	}

	if user.Password != defaults.String {
		updateFields["password"] = user.Password
	}

	if user.Email != defaults.String {
		updateFields["email"] = user.Email
	}

	if user.Name != defaults.String {
		updateFields["name"] = user.Name
	}

	err := r.db.Model(new(entity.User)).Where("id = ?", user.ID).Updates(updateFields).Error
	if err != nil {
		return nil, exception.Wrap(
			err,
			exception.ErrMySQLInternal,
			exception.StatusInternalServerError,
			"사용자를 수정할 수 없습니다.",
		)
	}

	return user, err
}

func (r UserRepository) DeleteUser(ctx context.Context, id int) error {
	err := r.db.WithContext(ctx).Delete(new(entity.User), id).Error
	if err != nil {
		return exception.Wrap(
			err,
			exception.ErrMySQLInternal,
			exception.StatusInternalServerError,
			"사용자를 삭제할 수 없습니다.",
		)
	}

	return nil
}

func (r UserRepository) GetLetters(ctx context.Context, userID int) ([]entity.Letter, error) {
	var foundUser entity.User
	err := r.db.WithContext(ctx).Model(&foundUser).Association("Letters").Error
	if err != nil {
		return nil, exception.Wrap(
			err,
			exception.ErrMySQLInternal,
			exception.StatusInternalServerError,
			"해당 사용자의 letters를 가져올 수 없습니다.",
			exception.WithData(
				exception.Map{
					"userID": userID,
				},
			),
		)
	}

	return foundUser.Letters, nil
}

func (r UserRepository) GetFeelings(ctx context.Context, userID int) ([]entity.Feeling, error) {
	var foundUser entity.User
	err := r.db.WithContext(ctx).Model(&foundUser).Association("Feelings").Error
	if err != nil {
		return nil, exception.Wrap(
			err,
			exception.ErrMySQLInternal,
			exception.StatusInternalServerError,
			"해당 사용자의 feelings 가져올 수 없습니다.",
			exception.WithData(
				exception.Map{
					"userID": userID,
				},
			),
		)
	}

	return foundUser.Feelings, nil
}
