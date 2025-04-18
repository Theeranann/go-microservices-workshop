package game

import (
	"encoding/json"
	"fmt"
	"game_service/internal/domain/entity"
	"game_service/pkg/logs"
	"io"
	"net/http"

	"gorm.io/gorm"
)

type gamesRepository struct {
	db *gorm.DB
}

func NewGamesRepositoryDB(db *gorm.DB) entity.GamesRepository {
	db.AutoMigrate(&entity.Game{})
	getGamesData(db)
	return &gamesRepository{db: db}
}

func (r gamesRepository) GetAll() (games []entity.Game, err error) {
	err = r.db.Order("id").Limit(30).Find(&games).Error
	if err != nil {
		logs.Errorf("Error fetching all games: %v", err)
	}
	return games, err
}

func (r gamesRepository) GetGameByID(gameID uint) (*entity.Game, error) {
	game := new(entity.Game)
	err := r.db.First(game, gameID).Error
	if err == gorm.ErrRecordNotFound {
		logs.Errorf("Game not found for ID: %d", gameID)
		return nil, nil
	} else if err != nil {
		logs.Errorf("Error fetching game by ID %d: %v", gameID, err)
		return nil, err
	}

	return game, err
}

func getGamesData(db *gorm.DB) error {
	// If not have games in database it will get data from API
	var count int64
	db.Model(&entity.Game{}).Count(&count)
	if count > 0 {
		return nil
	}
	apiURL := "https://www.freetogame.com/api/games"

	// Fetch data from the API
	var games []entity.Game
	if err := fetchAPI(apiURL, &games); err != nil {
		return err
	}

	// Insert data into the database
	for _, game := range games {
		result := db.Create(&game)
		if result.Error != nil {
			logs.Errorf("Error inserting record: %v", result.Error)
		}
	}

	logs.Info("Data inserted successfully.")
	return nil
}

func fetchAPI(apiURL string, target interface{}) error {
	// Make an HTTP GET request to the API
	response, err := http.Get(apiURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Check if the response status code is OK
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status code: %d", response.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// Unmarshal the JSON response into the target struct
	if err := json.Unmarshal(body, target); err != nil {
		return err
	}

	return nil
}
