package handlers

import (
	"matrix-api/internal/auth"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func HandleLogin(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if req.Username != "admin" || req.Password != "password123" {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	token, err := auth.GenerateToken(req.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error generating token")
	}

	return c.JSON(LoginResponse{Token: token})
}
