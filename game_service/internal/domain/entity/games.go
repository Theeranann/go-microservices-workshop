package entity

import (
	"game_service/internal/domain/event"
	"time"

	"gorm.io/gorm"
)

type GamesRepository interface {
	GetAll() ([]Game, error)
	GetGameByID(gameID uint) (*Game, error)
}

type GamesUsecase interface {
	GetAll() ([]Game, error)
	GetGameByID(gameID uint) (*Game, error)
	HandleUserRead(command event.UserReadCommand) error
}

// Game structure data from API
type Game struct {
	gorm.Model
	Title             string
	Thumbnail         string `gorm:"column:thumbnail"`
	ShortDescription  string `gorm:"column:short_description"`
	GameURL           string `gorm:"column:game_url"`
	Genre             string
	Platform          string
	Publisher         string
	Developer         string
	ReleaseDate       time.Time `gorm:"column:release_date"`
	FreeToGameProfile string    `gorm:"column:freetogame_profile_url"`
}

type APIData struct {
	Title             string `json:"title"`
	Thumbnail         string `json:"thumbnail"`
	ShortDescription  string `json:"short_description"`
	GameURL           string `json:"game_url"`
	Genre             string `json:"genre"`
	Platform          string `json:"platform"`
	Publisher         string `json:"publisher"`
	Developer         string `json:"developer"`
	ReleaseDate       string `json:"release_date"`
	FreeToGameProfile string `json:"freetogame_profile_url"`
}
