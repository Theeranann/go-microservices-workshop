package event

var Topics = []string{
	TopicUserCreated,
	TopicUserDeleted,
}

const (
	TopicUserCreated = "user.created"
	TopicUserDeleted = "user.deleted"
	TopicUserReaded  = "user.readed"
)

var TopicEventTypeMap = map[string]func() Event{
	TopicUserCreated: func() Event { return &UserCreatedEvent{} },
	TopicUserDeleted: func() Event { return &UserDeletedEvent{} },
}

type Event interface{}

type UserEventRepository interface {
	SaveUser(user UserCreatedEvent) error
	RemoveUser(user UserDeletedEvent) error
}

type UserEventUsecase interface {
	HandleUserCreated(user UserCreatedEvent) error
	HandleUserDeleted(user UserDeletedEvent) error
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

type UserCreatedEvent struct {
	UserID uint64 `json:"user_id" db:"user_id"`
}

type UserDeletedEvent struct {
	UserID uint64 `json:"user_id" db:"user_id"`
}
