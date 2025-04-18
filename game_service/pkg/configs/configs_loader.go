package configs

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

func LoadConfigs(cfg *Configs) {
	// Fiber configs
	cfg.App.Host = os.Getenv("FIBER_HOST")
	cfg.App.Port = os.Getenv("FIBER_PORT")
	if cfg.App.Port == "" {
		cfg.App.Port = "80" // default port
	}

	log.Println("Fiber Configs", logrus.Fields{
		"host": cfg.App.Host,
		"port": cfg.App.Port,
	})

	// Database Configs
	cfg.PostgreSQL.Host = os.Getenv("DB_HOST")
	cfg.PostgreSQL.Port = os.Getenv("DB_PORT")
	cfg.PostgreSQL.Protocol = os.Getenv("DB_PROTOCOL")
	cfg.PostgreSQL.Username = os.Getenv("DB_USERNAME")
	cfg.PostgreSQL.Password = os.Getenv("DB_PASSWORD")
	cfg.PostgreSQL.Database = os.Getenv("DB_DATABASE")

	// Redis Configs
	cfg.Redis.Host = os.Getenv("REDIS_HOST")
	cfg.Redis.Port = os.Getenv("REDIS_PORT")
	cfg.Redis.Password = os.Getenv("REDIS_PASSWORD")

	log.Println("Database Configs", logrus.Fields{
		"host":       cfg.PostgreSQL.Host,
		"port":       cfg.PostgreSQL.Port,
		"protocol":   cfg.PostgreSQL.Protocol,
		"username":   cfg.PostgreSQL.Username,
		"database":   cfg.PostgreSQL.Database,
		"redis_Host": cfg.Redis.Host,
		"redis_Port": cfg.Redis.Port,
	})
}
