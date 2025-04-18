package event

type GamesProducer interface {
	UserReaded(command ReadedEvent) error
}

type UserReadCommand struct {
	UserID string
	GameID string
}
