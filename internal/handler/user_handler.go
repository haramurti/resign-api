package handler

import (
	"resign-api/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	usecase domain.UserUsecase
}

func NewUserHandler(u domain.UserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
}

// getprofile hanlder
func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	userLocal := c.Locals("currentUser")
	if userLocal == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Identitas karyawan tidak ditemukan"})
	}

	user := userLocal.(domain.User)
	return c.JSON(user)
}
