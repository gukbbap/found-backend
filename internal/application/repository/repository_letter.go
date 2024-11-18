package repository

import "found-backend/internal/infra/storage/mysql"

type LetterRepository struct {
	mysql *mysql.MySQL
}

func NewLetterRepository(mysql *mysql.MySQL) *LetterRepository {
	return &LetterRepository{
		mysql: mysql,
	}
}
