package mysql

import (
	"fmt"
	"found-backend/internal/infra/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	*gorm.DB
}

func NewMySQL(config *config.Config) (*MySQL, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Schema,
	)

	openedDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		return nil, err
	}

	return &MySQL{
		DB: openedDB,
	}, nil
}
