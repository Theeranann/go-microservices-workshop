package servers

import (
	"context"
	"log"
	"time"
	"user_service/internal/domain/event"
	"user_service/pkg/configs"
	"user_service/pkg/logs"
	"user_service/pkg/utils"

	"gorm.io/gorm"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App             *fiber.App
	Cfg             *configs.Configs
	Db              *gorm.DB
	Consumer        sarama.ConsumerGroup
	ConsumerHandler sarama.ConsumerGroupHandler
}

func NewServer(cfg *configs.Configs, db *gorm.DB, consumer sarama.ConsumerGroup) *Server {
	return &Server{
		App:      fiber.New(),
		Cfg:      cfg,
		Db:       db,
		Consumer: consumer,
	}
}

func (s *Server) Start() {
	if err := s.MapHandlers(); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}

	fiberConnURL, err := utils.ConnectionUrlBuilder("fiber", s.Cfg)
	if err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}

	logs.Infof("Kafka consumer starting with topics: %v", event.Topics)
	go func() {
		log.Println("consumer start..")
		for {
			err := s.Consumer.Consume(context.Background(), event.Topics, s.ConsumerHandler)
			if err != nil {
				logs.Errorf("Kafka consumer error: %v", err)
				time.Sleep(5 * time.Second)
			}
		}
	}()

	host := s.Cfg.App.Host
	port := s.Cfg.App.Port
	log.Printf("server has been started on %s:%s ⚡", host, port)

	if err := s.App.Listen(fiberConnURL); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}
}
