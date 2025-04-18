package event

type UserProducer interface {
	UserCreate(command UserEventCommand) error
	UserDelete(command UserEventCommand) error
}

type UserEventCommand struct {
	UserID uint `json:"user_id"`
}
