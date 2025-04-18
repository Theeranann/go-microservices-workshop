package event

var Topics = []string{
	TopicUserReaded,
}

const (
	TopicUserCreated = "user.created"
	TopicUserDeleted = "user.deleted"
	TopicUserReaded  = "user.readed"
)

var TopicEventTypeMap = map[string]func() Event{
	TopicUserReaded: func() Event { return &ReadedEvent{} },
}

type Event interface {}

type GameEventRepository interface {
	SaveRead(req ReadedEvent) error
}

type GameEventUsecase interface {
	HandleUserRead(req ReadedEvent) error
}

type ReadedEvent struct {
	UserID    string
	GameID    string
	Title     string
	Genre     string
	Platform  string
	Publisher string
	Developer string
}

type CreateUserEvent struct {
	UserID uint64 `json:"user_id" db:"user_id"`
}

type DeleteUserEvent struct {
	UserID uint64 `json:"user_id" db:"user_id"`
}

// Topic implements Event.
func (d DeleteUserEvent) Topic() string {
	panic("unimplemented")
}
