package mysql

import "found-backend/internal/application/entity"

func (mysql MySQL) AutoMigration() error {
	return mysql.db.AutoMigrate(
		new(entity.User),
		new(entity.Letter),
		new(entity.Feeling),
	)
}
