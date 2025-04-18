package response

import (
	"github.com/gofiber/fiber/v2"
)

type BaseResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Result     interface{} `json:"result,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}

// Success response with data
func Success(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(BaseResponse{
		StatusCode: fiber.StatusOK,
		Message:    message,
		Result:     data,
	})
}

// Created response
func Created(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(BaseResponse{
		StatusCode: fiber.StatusCreated,
		Message:    message,
		Result:     data,
	})
}

// Error response
func Error(c *fiber.Ctx, statusCode int, message string, errDetail interface{}) error {
	return c.Status(statusCode).JSON(BaseResponse{
		StatusCode: statusCode,
		Message:    message,
		Error:      errDetail,
	})
}
