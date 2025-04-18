package game

import (
	"fmt"
	"game_service/internal/domain/entity"
	"game_service/internal/domain/event"
	"game_service/pkg/logs"
	"strconv"
)

type GamesUsecase struct {
	GamesRepo     entity.GamesRepository
	GameRedis     entity.GameRepositoryRedis
	EventProducer event.GamesProducer
}

// Constructor
func NewGamesUsecase(GamesRepo entity.GamesRepository,
	redis entity.GameRepositoryRedis,
	producer event.GamesProducer,
) entity.GamesUsecase {
	return &GamesUsecase{
		GamesRepo:     GamesRepo,
		GameRedis:     redis,
		EventProducer: producer,
	}
}

func (u *GamesUsecase) GetAll() (games []entity.Game, err error) {
	// Try to get data from Redis
	cachedGames, err := u.GameRedis.GetGameListFromCache()
	if err == nil && cachedGames != nil {
		logs.Info("[GET games/] Returning data from Redis cache")
		return cachedGames, nil
	}

	logs.Info("[GET games/] Cache miss -> fetching from DB")

	// If Redis failed (Cache Miss), get data from the database
	games, err = u.GamesRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch games from DB: %w", err)
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				logs.Errorf("panic while caching to Redis: %v", r)
			}
		}()
		err := u.GameRedis.CacheGameList(games)
		if err != nil {
			logs.Errorf("[GET games/] Failed to cache games: %v", err)
		}
	}()

	logs.Info("[GET games/] Game list cached to Redis")
	return games, err
}

func (u *GamesUsecase) GetGameByID(gameID uint) (*entity.Game, error) {
	// Try to get data from Redis
	cachedGames, err := u.GameRedis.GetGameByIdFromCache(gameID)
	if err == nil && cachedGames != nil {
		logs.Info("[GET games/:id] Returning data from Redis cache")
		return cachedGames, nil
	}

	logs.Info("[GET games/:id] Cache miss -> fetching from DB")

	// If Redis failed (Cache Miss), get data from the database
	game, err := u.GamesRepo.GetGameByID(gameID)
	if err != nil {
		return nil, err
	}

	// Cache the result in Redis
	go func() {
		defer func() {
			if r := recover(); r != nil {
				logs.Errorf("panic while caching to Redis: %v", r)
			}
		}()
		err := u.GameRedis.CacheGameById(gameID, game)
		if err != nil {
			logs.Errorf("[GET games/:id] Failed to cache game: %v", err)
		}
	}()

	logs.Info("[GET games/:id] Game cached to Redis")
	return game, nil
}

func (u *GamesUsecase) HandleUserRead(command event.UserReadCommand) error {
	// Validation
	if command.GameID == "" || command.UserID == "" {
		logs.Infof("Missing required fields - GameID: %q, UserID: %q", command.GameID, command.UserID)
		return fmt.Errorf("bad request: missing GameID or UserID")
	}

	// Parse GameID
	gameID, err := strconv.ParseUint(command.GameID, 10, 64)
	if err != nil {
		logs.Errorf("Invalid GameID: %v", err)
		return fmt.Errorf("invalid GameID format: %w", err)
	}

	// Fetch game
	game, err := u.GamesRepo.GetGameByID(uint(gameID))
	if err != nil {
		logs.Errorf("Error fetching game by ID %d: %v", gameID, err)
		return fmt.Errorf("failed to get game: %w", err)
	}

	// Prepare event payload
	eventPayload := event.ReadedEvent{
		UserID:    command.UserID,
		GameID:    command.GameID,
		Title:     game.Title,
		Genre:     game.Genre,
		Platform:  game.Platform,
		Publisher: game.Publisher,
		Developer: game.Developer,
	}

	// Produce event
	err = u.EventProducer.UserReaded(eventPayload)
	if err != nil {
		logs.Errorf("Error producing UserReadedEvent: %v", err)
		return fmt.Errorf("failed to produce event: %w", err)
	}

	logs.Infof("[POST games/:id] %s event produced successfully for user ID: %s and game: %s", event.TopicUserReaded, command.UserID, game.Title)
	return nil
}
