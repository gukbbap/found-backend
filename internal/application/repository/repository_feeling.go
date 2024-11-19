package repository

import (
	"context"
	"found-backend/internal/application/entity"
	"found-backend/internal/infra/exception"
	"found-backend/internal/infra/storage/mysql"
)

type FeelingRepository struct {
	db *mysql.MySQL
}

func NewFeelingRepository(db *mysql.MySQL) *FeelingRepository {
	return &FeelingRepository{
		db: db,
	}
}

func (r FeelingRepository) CreateFeeling(ctx context.Context, feeling *entity.Feeling) (*entity.Feeling, error) {
	err := r.db.WithContext(ctx).Create(feeling).Error
	if err != nil {
		return nil, exception.Wrap(
			err,
			exception.ErrMySQLInternal,
			exception.StatusInternalServerError,
			"feeling을 생성할 수 없습니다.",
		)
	}

	return feeling, nil
}

func (r FeelingRepository) FindFeeling(ctx context.Context, id int) (*entity.Feeling, error) {
	var foundFeeling entity.Feeling
	err := r.db.WithContext(ctx).First(&foundFeeling, id).Error
	if err != nil {
		return nil, exception.Wrap(
			err,
			exception.ErrMySQLNotFound,
			exception.StatusNotFound,
			"feeling을 조회할 수 없습니다.",
		)
	}

	return &foundFeeling, nil
}

func (r FeelingRepository) UpdateFeeling(ctx context.Context, feeling *entity.Feeling) (*entity.Feeling, error) {
	err := r.db.WithContext(ctx).Save(feeling).Error
	if err != nil {
		return nil, exception.Wrap(
			err,
			exception.ErrMySQLInternal,
			exception.StatusInternalServerError,
			"feeling을 수정할 수 없습니다.",
		)
	}

	return feeling, nil
}

func (r FeelingRepository) DeleteFeeling(ctx context.Context, id int) error {
	err := r.db.WithContext(ctx).Delete(new(entity.Feeling), id).Error
	if err != nil {
		return exception.Wrap(
			err,
			exception.ErrMySQLInternal,
			exception.StatusInternalServerError,
			"feeling을 삭제할 수 없습니다.",
		)
	}

	return nil
}
