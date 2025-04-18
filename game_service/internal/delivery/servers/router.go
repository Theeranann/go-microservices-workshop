package servers

import (
	"fmt"
	_gameHttp "game_service/internal/delivery/handlers/game"
	_gamesHandlerProducer "game_service/internal/event/producer"
	_gameRepository "game_service/internal/repository/postgres/game"
	_gameRepositoryRedis "game_service/internal/repository/redis/game"
	_gamesUsecase "game_service/internal/usecase/game"

	_usersEventHandler "game_service/internal/event/consumer"
	_userEventRepository "game_service/internal/repository/postgres/user_event"
	_userEventUsecase "game_service/internal/usecase/user_event"

	_kafka "game_service/pkg/kafka"
	"game_service/pkg/logs"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/redis/go-redis/v9"
)

func (s *Server) MapHandlers() error {
	//middleware cors allow
	s.App.Use(cors.New(cors.Config{
		AllowOrigins: s.Cfg.App.Cors,
		AllowMethods: "GET, POST, PUT, DELETE, PATCH, OPTIONS",
		AllowHeaders: "Content-Type, Authorization",
	}))

	// Initialize Kafka producer
	kafkaProducer, err := initKafkaProducer()
	if err != nil {
		logs.Fatalf("Error initializing Kafka producer: %v", err)
		return err
	}

	// Initialize Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", s.Cfg.Redis.Host, s.Cfg.Redis.Port),
		Password: s.Cfg.Redis.Password,
	})

	// Initialize EventProducer with Kafka producer
	eventProducer := _kafka.NewEventProducer(kafkaProducer)
	gameProducer := _gamesHandlerProducer.NewGameProducer(eventProducer)

	// Initialize redisClient with RepositoryRedis
	gamesRepositoryRedis := _gameRepositoryRedis.NewGameRepositoryRedis(redisClient)

	// Group a version
	v1 := s.App.Group("/v1")

	//* Games group
	gamesGroup := v1.Group("/games")
	gameRepository := _gameRepository.NewGamesRepositoryDB(s.Db)
	gameUsecase := _gamesUsecase.NewGamesUsecase(gameRepository, gamesRepositoryRedis, gameProducer)
	_gameHttp.NewGamesHandler(gamesGroup, gameUsecase)

	//* User Event group
	userEventRepository := _userEventRepository.NewUserEventRepository(s.Db)
	userEventUsecase := _userEventUsecase.NewUserEventUse(userEventRepository)

	// Kafka consumer
	userEventHandler := _usersEventHandler.NewUserEventHandler(userEventUsecase)
	userConsumerHandler := _kafka.NewConsumerHandler(userEventHandler)
	s.ConsumerHandler = userConsumerHandler

	// End point not found response
	s.App.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     "error, end point not found",
			"result":      nil,
		})
	})

	return nil
}

func initKafkaProducer() (sarama.SyncProducer, error) {
	// Initialize Kafka producer
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		return nil, err
	}

	return producer, nil
}
