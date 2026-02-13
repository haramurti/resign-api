package handler

import (
	"resign-api/internal/domain"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	usecase domain.UserUsecase
}

func NewUserHandler(u domain.UserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
}

// Register: Handler buat daftar karyawan baru
func (h *UserHandler) Register(c *fiber.Ctx) error {
	var user domain.User

	// 1. Parsing JSON body ke struct User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Format data salah"})
	}

	// 2. Ambil context dari Fiber dan oper ke Usecase
	ctx := c.UserContext()
	if err := h.usecase.Register(ctx, &user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "User berhasil didaftarkan",
		"data":    user,
	})
}

// GetProfile: Handler buat ambil data profil
func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	// Ambil ID dari parameter URL /users/:id
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	ctx := c.UserContext()
	user, err := h.usecase.GetProfile(ctx, uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User tidak ditemukan"})
	}

	return c.JSON(user)
}
