package repository

import (
	"found-backend/internal/infra/storage/mysql"
)

type UserRepository struct {
	*mysql.MySQL
}

func NewUserRepository(mysql *mysql.MySQL) *UserRepository {
	return &UserRepository{
		MySQL: mysql,
	}
}
