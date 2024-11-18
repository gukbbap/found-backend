package repository

import "found-backend/internal/infra/storage/mysql"

type FeelingRepository struct {
	db *mysql.MySQL
}

func newFeelingRepository(db *mysql.MySQL) *FeelingRepository {
	return &FeelingRepository{
		db: db,
	}
}
