package game

import (
	"context"
	"encoding/json"
	"fmt"
	"game_service/internal/domain/entity"
	"game_service/pkg/logs"
	"time"

	"github.com/redis/go-redis/v9"
)

type GameRepositoryRedis struct {
	Redis *redis.Client
}

func NewGameRepositoryRedis(redis *redis.Client) entity.GameRepositoryRedis {
	return &GameRepositoryRedis{Redis: redis}
}

const redisTTL = time.Second * 120

func (r *GameRepositoryRedis) CacheGameById(gameId uint, game *entity.Game) error {
	key := fmt.Sprintf("game:id:%d", gameId)

	gameJSON, err := json.Marshal(game)
	if err != nil {
		logs.Errorf("[Redis] CacheGame failed to marshal: %v", err)
		return err
	}

	err = r.Redis.Set(context.Background(), key, gameJSON, redisTTL).Err()
	if err != nil {
		logs.Error("error set CacheGameById to redis")
		return err
	}

	return nil
}

func (r *GameRepositoryRedis) CacheGameList(games []entity.Game) error {
	key := "game:list"

	gamesListJSON, err := json.Marshal(games)
	if err != nil {
		logs.Errorf("[Redis] CacheGameList failed to marshal: %v", err)
		return err
	}

	err = r.Redis.Set(context.Background(), key, gamesListJSON, redisTTL).Err()
	if err != nil {
		logs.Error("error set CacheGameList to redis")
		return err
	}

	return nil
}

func (r *GameRepositoryRedis) GetGameByIdFromCache(gameId uint) (*entity.Game, error) {
	key := fmt.Sprintf("game:id:%d", gameId)

	result, err := r.Redis.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		logs.Errorf("error getting games from redis: %v", err)
		return nil, err
	}

	var games entity.Game
	if result != "" {
		err = json.Unmarshal([]byte(result), &games)
		if err != nil {
			logs.Errorf("error Unmarshal games: %v", err)
			return nil, err
		}
		return &games, nil
	}
	return nil, nil
}

func (r *GameRepositoryRedis) GetGameListFromCache() ([]entity.Game, error) {
	key := "game:list"

	result, err := r.Redis.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		logs.Errorf("error getting games from redis: %v", err)
		return nil, err
	}

	var games []entity.Game
	if result != "" {
		err = json.Unmarshal([]byte(result), &games)
		if err != nil {
			logs.Errorf("error Unmarshal games: %v", err)
			return nil, err
		}
		return games, nil
	}

	return nil, nil
}
