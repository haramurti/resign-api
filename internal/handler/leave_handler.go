package handler

import (
	"resign-api/internal/domain"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type LeaveHandler struct {
	usecase domain.LeaveUsecase
}

func NewLeaveHandler(u domain.LeaveUsecase) *LeaveHandler {
	return &LeaveHandler{usecase: u}
}

// Apply: User mengajukan cuti
func (h *LeaveHandler) Apply(c *fiber.Ctx) error {
	var leave domain.LeaveRequest
	if err := c.BodyParser(&leave); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Data input tidak valid"})
	}

	ctx := c.UserContext()
	if err := h.usecase.Apply(ctx, &leave); err != nil {
		// Pesan error dari usecase (misal: jatah habis) bakal dikirim ke sini
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Pengajuan cuti berhasil dikirim", "data": leave})
}

// GetHistory: Ambil semua daftar cuti (buat HR)
func (h *LeaveHandler) GetHistory(c *fiber.Ctx) error {
	ctx := c.UserContext()
	history, err := h.usecase.GetHistory(ctx)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data"})
	}
	return c.JSON(history)
}

// Approve: HR menyetujui cuti
func (h *LeaveHandler) Approve(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	ctx := c.UserContext()
	if err := h.usecase.ApproveLeave(ctx, uint(id)); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Cuti disetujui, jatah cuti user otomatis berkurang"})
}
