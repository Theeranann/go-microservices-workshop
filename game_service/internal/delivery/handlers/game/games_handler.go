package game

import (
	"game_service/internal/delivery/handlers/response"
	"game_service/internal/domain/entity"
	"game_service/internal/domain/event"
	"game_service/pkg/logs"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type gamesHandler struct {
	GamesUse entity.GamesUsecase
}

func NewGamesHandler(r fiber.Router, gameUse entity.GamesUsecase) {
	handler := &gamesHandler{GamesUse: gameUse}

	r.Get("/", handler.GetAllGames)
	r.Get("/:id", handler.GetGameByID)
	r.Post("/", handler.UserRead)
}

func (h *gamesHandler) GetAllGames(c *fiber.Ctx) error {
	logs.Info("[GET games/] Request received")

	games, err := h.GamesUse.GetAll()
	if err != nil {
		logs.Errorf("Error getting all games: %v", err)
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get games", err.Error())
	}

	logs.Info("[GET games/] Response sent successfully")
	return response.Success(c, "Games retrieved successfully", games)
}

func (h *gamesHandler) GetGameByID(c *fiber.Ctx) error {
	logs.Info("[GET games/:id] Request received")

	gameIDParam := c.Params("id")
	gameID, err := strconv.ParseUint(gameIDParam, 10, 64)
	if err != nil {
		logs.Errorf("Invalid game ID: %v", err)
		return response.Error(c, fiber.StatusBadRequest, "Invalid game ID", err.Error())
	}

	game, err := h.GamesUse.GetGameByID(uint(gameID))
	if err != nil {
		logs.Errorf("Error getting game by ID %d: %v", gameID, err)
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get game by ID", err.Error())
	}

	if game == nil {
		return response.Error(c, fiber.StatusNotFound, "Game not found", nil)
	}

	logs.Info("[GET games/:id] Response sent successfully")
	return response.Success(c, "Game retrieved successfully", game)
}

func (h *gamesHandler) UserRead(c *fiber.Ctx) error {
	logs.Info("[POST games/] Request received")

	command := event.UserReadCommand{}
	err := c.BodyParser(&command)
	if err != nil {
		logs.Errorf("Error parsing UserRead command: %v", err)
		return response.Error(c, fiber.StatusBadRequest, "Invalid payload", err.Error())
	}

	logs.Debugf("[POST games/:id] Payload %+v", command)

	err = h.GamesUse.HandleUserRead(command)
	if err != nil {
		logs.Errorf("Error producing UserRead event: %v", err)
		return response.Error(c, fiber.StatusInternalServerError, "Failed to process UserRead event", err.Error())
	}

	logs.Info("[POST games/:id] Response sent successfully")
	return response.Created(c, "UserRead success", nil)
}
