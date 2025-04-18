package entity

type GameRepositoryRedis interface {
	// Get data from Redis
	GetGameListFromCache() ([]Game, error)
	GetGameByIdFromCache(gameId uint) (*Game, error)

	// Set data to Redis
	CacheGameList(games []Game) error
	CacheGameById(gameId uint, games *Game) error
}