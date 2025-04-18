package gameevent

import (
	"user_service/internal/domain/event"
	"user_service/pkg/logs"

	"gorm.io/gorm"
)

type gameEventRepository struct {
	DB *gorm.DB
}

func NewGameEventRepository(Db *gorm.DB) event.GameEventRepository {
	Db.AutoMigrate(&event.ReadedEvent{})
	return &gameEventRepository{DB: Db}
}

func (r *gameEventRepository) SaveRead(req event.ReadedEvent) error {
	err := r.DB.Create(&req).Error
	if err != nil {
		logs.Errorf("Error creating read event: %v", err)
		return err
	}
	return nil
}