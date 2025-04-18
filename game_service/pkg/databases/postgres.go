package database

import (
	"game_service/pkg/configs"
	"game_service/pkg/logs"
	"game_service/pkg/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgreSQLDBConnection(cfg *configs.Configs) (*gorm.DB, error) {
	postgresUrl, err := utils.ConnectionUrlBuilder("postgresql", cfg)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(postgresUrl), &gorm.Config{})
	if err != nil {
		logs.Errorf("error, can't connect to database, %s", err.Error())
		return nil, err
	}

	logs.Info("postgreSQL database has been connected 🐘")
	return db, nil
}
