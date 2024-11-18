package repository

import "found-backend/internal/infra/storage/mysql"

type LetterRepository struct {
	db *mysql.MySQL
}

func NewLetterRepository(db *mysql.MySQL) *LetterRepository {
	return &LetterRepository{
		db: db,
	}
}
