package auth

import (
	"user_service/internal/delivery/handlers/response"
	middlewares "user_service/internal/delivery/middleware"
	"user_service/internal/domain/entity"
	"user_service/pkg/configs"
	"user_service/pkg/logs"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	Cfg     *configs.Configs
	AuthUse entity.AuthUsecase
}

func NewAuthHandler(r fiber.Router, cfg *configs.Configs, authUse entity.AuthUsecase) {
	Handler := &authHandler{
		Cfg:     cfg,
		AuthUse: authUse,
	}

	r.Post("/login", Handler.Login)
	r.Get("/auth-test", middlewares.JwtAuthentication(), Handler.AuthTest)
}

func (h *authHandler) Login(c *fiber.Ctx) error {
	logs.Info("[POST auth/login] Request received")

	req := new(entity.UsersCredentials)
	if err := c.BodyParser(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error(), err)
	}

	res, err := h.AuthUse.Login(h.Cfg, req)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error(), err)
	}

	logs.Infof("[POST auth/login] User login successfully")
	return response.Success(c, "User login successfully", res)
}

func (h *authHandler) AuthTest(c *fiber.Ctx) error {
	logs.Info("[GET auth/auth-test] Request received")
	id := c.Locals("user_id")
	username := c.Locals("username")

	logs.Info("[GET auth/auth-test] Auth-Test successfully")
	return response.Success(c, "AuthTest", fiber.Map{"id": id, "username": username})
}
