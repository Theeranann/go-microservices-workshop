package main

import (
	"fmt"
	"game_service/internal/delivery/servers"
	"game_service/pkg/configs"
	database "game_service/pkg/databases"
	"game_service/pkg/logs"
	"github.com/IBM/sarama"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	// Load dotenv config
	err := godotenv.Load(".env")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Error: The .env file does not exist.")
		} else {
			fmt.Println("Error loading .env file:", err)
		}
		panic(err.Error())
	}

	cfg := new(configs.Configs)
	logs.Info("Loading configuration...")
	configs.LoadConfigs(cfg)
	logs.Info("Configuration loaded successfully.")

	// Kafka Configs
	kafkaBroker := os.Getenv("KAFKA_BROKER")

	// New Database
	db, err := database.NewPostgreSQLDBConnection(cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer func(db *gorm.DB) {
		sqlDb, err := db.DB()
		if err != nil {
			logs.Errorf("Error getting SQL database: %v", err)
		}
		if err := sqlDb.Close(); err != nil {
			logs.Errorf("Error closing database connection: %v", err)
		}
	}(db)

	// Create a Kafka consumer
	consumer, err := sarama.NewConsumerGroup([]string{kafkaBroker}, "group_id_service_game", nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	s := servers.NewServer(cfg, db, consumer)
	s.Start()
}
