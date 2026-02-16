package handler

import (
	"resign-api/internal/domain"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type LeaveHandler struct {
	usecase domain.LeaveUsecase
}

func NewLeaveHandler(u domain.LeaveUsecase) *LeaveHandler {
	return &LeaveHandler{usecase: u}
}

// apply leave
func (h *LeaveHandler) Apply(c *fiber.Ctx) error {
	// Kita pake struct temporary biar gampang parsing tanggal dari string JSON
	var input struct {
		UserID    uint   `json:"user_id"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
		Reason    string `json:"reason"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Format input salah"})
	}

	//conversion
	start, _ := time.Parse("2006-01-02", input.StartDate)
	end, _ := time.Parse("2006-01-02", input.EndDate)

	leave := domain.LeaveRequest{
		UserID:    input.UserID,
		StartDate: start,
		EndDate:   end,
		Reason:    input.Reason,
		Status:    "pending",
	}

	ctx := c.UserContext()
	if err := h.usecase.Apply(ctx, &leave); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Cuti dari " + input.StartDate + " sampai " + input.EndDate + " berhasil diajukan",
		"data":    leave,
	})
}

func (h *LeaveHandler) GetHistory(c *fiber.Ctx) error {
	ctx := c.UserContext()
	history, err := h.usecase.GetHistory(ctx)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data"})
	}
	return c.JSON(history)
}

func (h *LeaveHandler) Approve(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	ctx := c.UserContext()
	if err := h.usecase.ApproveLeave(ctx, uint(id)); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Cuti disetujui, jatah cuti user otomatis berkurang"})
}
