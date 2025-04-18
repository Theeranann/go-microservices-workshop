package user

import (
	"strconv"
	"user_service/internal/delivery/handlers/response"
	"user_service/internal/domain/entity"
	"user_service/pkg/logs"

	"github.com/gofiber/fiber/v2"
)

type usersHandler struct {
	UsersUse entity.UsersUsecase
}

func NewUsersHandler(r fiber.Router, usersUse entity.UsersUsecase) {
	handler := &usersHandler{UsersUse: usersUse}

	r.Post("/", handler.Register)
	r.Get("/:id", handler.GetUserByID)
	r.Put("/:id", handler.UpdateUser)
	r.Delete("/:id", handler.DeleteUser)
	r.Get("/:id/read", handler.GetReadHistory)
}

func (h *usersHandler) Register(c *fiber.Ctx) error {
	logs.Info("[POST users/] Request received")

	req := new(entity.UsersRegisterReq)
	if err := c.BodyParser(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error(), err.Error())
	}

	res, err := h.UsersUse.Register(req)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to Register", err.Error())
	}

	logs.Infof("[POST users/] User register successfully")
	return response.Success(c, "Register successfully", res)
}

func (h *usersHandler) GetUserByID(c *fiber.Ctx) error {
	logs.Info("[GET users/:id] Request received")

	// Get the user ID from the URL parameters
	userIDParam := c.Params("id")
	userID, err := strconv.ParseUint(userIDParam, 10, 64)

	if err != nil {
		logs.Errorf("Error parsing user ID from URL parameters: %v", err)
		return response.Error(c, fiber.StatusBadRequest, "Error parsing user ID from URL parameters", err.Error())
	}

	// Call the GetUserByID method from the use case
	user, err := h.UsersUse.GetUserByID(uint(userID))
	if err != nil {
		logs.Errorf("Error getting user by ID (%d): %v", userID, err)
		return response.Error(c, fiber.StatusInternalServerError, "Error getting user by ID", err.Error())
	}

	if user == nil {
		logs.Errorf("User not found for ID: %d", userID)
		return response.Error(c, fiber.StatusNotFound, "User not found for ID", nil)
	}

	return response.Success(c, "User retrieved successfully", entity.UsersRes{Id: user.Id, Username: user.Username})
}

func (h *usersHandler) UpdateUser(c *fiber.Ctx) error {
	logs.Info("[PUT users/:id] Request received")

	// Get the user ID from the URL parameters
	userIDParam := c.Params("id")
	userID, err := strconv.ParseUint(userIDParam, 10, 64)

	if err != nil {
		logs.Errorf("Error parsing user ID from URL parameters: %v", err)
		return response.Error(c, fiber.StatusBadRequest, "Invalid user ID", err.Error())
	}

	// Parse the request body to get the updated user information
	req := new(entity.UsersUpdateReq)
	err = c.BodyParser(req)
	if err != nil {
		logs.Errorf("Error parsing update request body: %v", err)
		return response.Error(c, fiber.StatusBadRequest, "Error parsing update request body", err.Error())
	}

	// Call the UpdateUser method from the use case
	updatedUser, err := h.UsersUse.UpdateUser(uint(userID), req)
	if err != nil {
		logs.Errorf("Error updating user (ID=%d): %v", userID, err)
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update user", err.Error())
	}

	if updatedUser == nil {
		logs.Errorf("User not found for update: ID=%d", userID)
		return response.Error(c, fiber.StatusNotFound, "User not found", nil)
	}

	return response.Success(c, "User updated successfully", entity.UsersRes{Id: updatedUser.Id, Username: updatedUser.Username})
}

func (h *usersHandler) DeleteUser(c *fiber.Ctx) error {
	logs.Info("[DELETE users/:id] Request received")

	// Get the user ID from the URL parameters
	userIDParam := c.Params("id")
	userID, err := strconv.ParseUint(userIDParam, 10, 64)

	if err != nil {
		logs.Errorf("Error parsing user ID from URL parameters: %v", err)
		return response.Error(c, fiber.StatusBadRequest, "Invalid user ID", err.Error())
	}

	// Call the DeleteUser method from the use case
	err = h.UsersUse.DeleteUser(uint(userID))
	if err != nil {
		logs.Errorf("Error deleting user (ID=%d): %v", userID, err)
		return response.Error(c, fiber.StatusInternalServerError, "Failed to delete user", err.Error())
	}

	return response.Success(c, "User deleted successfully", nil)
}

func (h *usersHandler) GetReadHistory(c *fiber.Ctx) error {
	logs.Info("[GET users/:id/read] Request received")

	// Get the user ID from the URL parameters
	userIDParam := c.Params("id")
	userID, err := strconv.ParseUint(userIDParam, 10, 64)

	if err != nil {
		logs.Errorf("Error parsing user ID from URL parameters: %v", err)
		return response.Error(c, fiber.StatusBadRequest, "Invalid user ID", err.Error())
	}

	// Call the GetReadHistory method from the use case
	readed, err := h.UsersUse.GetReadHistory(uint(userID))
	if err != nil {
		logs.Errorf("Error getting Readed by userID (ID=%d): %v", userID, err)
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get Readed by user ID", err.Error())
	}

	if readed == nil {
		logs.Errorf("Readed not found for user ID: ID=%d", userID)
		return response.Error(c, fiber.StatusNotFound, "Readed not found for user ID", nil)
	}

	return response.Success(c, "UserReaded retrieved successfully", readed)
}
