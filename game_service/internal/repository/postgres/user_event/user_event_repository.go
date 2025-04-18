package userevent

import (
	"game_service/internal/domain/event"
	"game_service/pkg/logs"

	"gorm.io/gorm"
)

type userEventRepository struct {
	Db *gorm.DB
}

func NewUserEventRepository(Db *gorm.DB) event.UserEventRepository {
	Db.AutoMigrate(&event.UserCreatedEvent{})
	return &userEventRepository{Db: Db}
}

func (r *userEventRepository) SaveUser(user event.UserCreatedEvent) error {
	err := r.Db.Create(&user).Error
	if err != nil {
		logs.Errorf("Error creating user event: %v", err)
		return err
	}

	return nil
}

func (r *userEventRepository) RemoveUser(user event.UserDeletedEvent) error {
	result := r.Db.Where("user_id = ?", user.UserID).Delete(&event.UserCreatedEvent{})
	if result.Error != nil {
		logs.Errorf("Error deleting user event: %v", result.Error)
		return result.Error
	}

	return nil
}
