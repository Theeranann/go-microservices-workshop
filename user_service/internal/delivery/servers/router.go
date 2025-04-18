package servers

import (
	_kafka "user_service/pkg/kafka"
	"user_service/pkg/logs"

	_usersHttp "user_service/internal/delivery/handlers/user"
	_usersEventHandler "user_service/internal/event/consumer"
	_usersHandlerProducer "user_service/internal/event/producer"
	_usersRepository "user_service/internal/repository/postgres/user"
	_usersUsecase "user_service/internal/usecase/user"

	_authHttp "user_service/internal/delivery/handlers/auth"
	_authRepository "user_service/internal/repository/postgres/auth"
	_authUsecase "user_service/internal/usecase/auth"

	_gameEventRepository "user_service/internal/repository/postgres/game_event"
	_gameEventUsecase "user_service/internal/usecase/game_event"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	// Group a version
	v1 := s.App.Group("/v1")

	// Initialize EventProducer with Kafka producer
	eventProducer := _kafka.NewEventProducer(kafkaProducer)
	userProducer := _usersHandlerProducer.NewUserProducer(eventProducer)

	//* Users group
	usersGroup := v1.Group("/users")
	usersRepository := _usersRepository.NewUsersRepositoryDB(s.Db)
	usersUsecase := _usersUsecase.NewUsersUseCase(usersRepository, userProducer)
	_usersHttp.NewUsersHandler(usersGroup, usersUsecase)

	//* Auth group
	authGroup := v1.Group("/auth")
	authRepository := _authRepository.NewAuthRepository(s.Db)
	authUsecase := _authUsecase.NewAuthUsecase(authRepository, usersRepository)
	_authHttp.NewAuthHandler(authGroup, s.Cfg, authUsecase)

	// Game Event group
	gameEventRepository := _gameEventRepository.NewGameEventRepository(s.Db)
	gameEventUsecase := _gameEventUsecase.NewGameEventUse(gameEventRepository)
	
	// Kafka consumer
	userEventHandler := _usersEventHandler.NewUserEventHandler(gameEventUsecase)
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
