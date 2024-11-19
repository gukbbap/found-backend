package repository

import (
	"context"
	"found-backend/internal/application/entity"
	"found-backend/internal/infra/exception"
	"found-backend/internal/infra/storage/mysql"
)

type LetterRepository struct {
	db *mysql.MySQL
}

func NewLetterRepository(db *mysql.MySQL) *LetterRepository {
	return &LetterRepository{
		db: db,
	}
}

func (r LetterRepository) CreateLetter(ctx context.Context, letter *entity.Letter) (*entity.Letter, error) {
	err := r.db.WithContext(ctx).Create(letter).Error
	if err != nil {
		return nil, exception.Wrap(
			err,
			exception.ErrMySQLInternal,
			exception.StatusInternalServerError,
			"글을 생성할 수 없습니다.",
		)
	}

	return letter, nil
}

func (r LetterRepository) FindLetter(ctx context.Context, id int) (*entity.Letter, error) {
	var foundLetter entity.Letter
	err := r.db.WithContext(ctx).First(&foundLetter, id).Error
	if err != nil {
		return nil, exception.Wrap(
			err,
			exception.ErrMySQLNotFound,
			exception.StatusNotFound,
			"글을 조회할 수 없습니다.",
		)
	}

	return &foundLetter, nil
}

func (r LetterRepository) UpdateLetter(ctx context.Context, letter *entity.Letter) (*entity.Letter, error) {
	err := r.db.WithContext(ctx).Save(letter).Error
	if err != nil {
		return nil, exception.New(
			exception.ErrMySQLInternal,
			exception.StatusInternalServerError,
			"글을 수정할 수 없습니다.",
		)
	}

	return letter, nil
}

func (r LetterRepository) DeleteLetter(ctx context.Context, id int) error {
	err := r.db.WithContext(ctx).Delete(new(entity.Letter), id).Error
	if err != nil {
		return exception.Wrap(
			err,
			exception.ErrMySQLInternal,
			exception.StatusInternalServerError,
			"글을 삭제할 수 없습니다.",
		)
	}

	return nil
}

func (r LetterRepository) GetFeelings(ctx context.Context, id int) ([]entity.Feeling, error) {
	var foundLetter entity.Letter

	err := r.db.WithContext(ctx).Model(&foundLetter).Association("feelings").Error
	if err != nil {
		return nil, exception.Wrap(
			err,
			exception.ErrMySQLInternal,
			exception.StatusInternalServerError,
			"해당 글의 feelings를 조회할 수 없습니다.",
			exception.WithData(
				exception.Map{
					"id": id,
				},
			),
		)
	}

	return foundLetter.Feelings, nil
}
