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

// GetProfile: Satu-satunya handler untuk identitas karyawan
func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	// 1. Ambil data dari Locals (hasil kerja keras NewAuthMiddleware lo)
	userLocal := c.Locals("currentUser")
	if userLocal == nil {
		// Proteksi: Jika entah bagaimana Locals kosong, langsung tolak
		return c.Status(401).JSON(fiber.Map{"error": "Identitas karyawan tidak ditemukan"})
	}

	// 2. Casting data ke struct User dan balikin sebagai JSON
	user := userLocal.(domain.User)
	return c.JSON(user)
}
