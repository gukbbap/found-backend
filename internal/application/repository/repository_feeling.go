package repository

import "found-backend/internal/infra/storage/mysql"

type FeelingRepository struct {
	mysql *mysql.MySQL
}

func newFeelingRepository(mysql *mysql.MySQL) *FeelingRepository {
	return &FeelingRepository{
		mysql: mysql,
	}
}
