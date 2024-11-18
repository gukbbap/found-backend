package main

import (
	"found-backend/internal/infra/config"
	"found-backend/internal/infra/storage/mysql"
	"log"
)

func main() {
	cfg, err := config.NewConfig(".env.toml")
	if err != nil {
		log.Panicf("failed to new config: %v\n", err)
	}

	mysql, err := mysql.NewMySQL(cfg)
	if err != nil {
		log.Panicf("failed to new mysql: %v\n", err)
	}

	err = mysql.AutoMigration()
	if err != nil {
		log.Panicf("failed to auto migration: %v\n", err)
	}
}
