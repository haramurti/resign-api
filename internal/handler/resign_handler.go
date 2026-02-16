package handler

import (
	"resign-api/internal/domain"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ResignationHandler struct {
	usecase domain.ResignationUsecase
}

func NewResignationHandler(u domain.ResignationUsecase) *ResignationHandler {
	return &ResignationHandler{usecase: u}
}

func (h *ResignationHandler) Submit(c *fiber.Ctx) error {
	var resign domain.Resignation
	if err := c.BodyParser(&resign); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
	}

	ctx := c.UserContext()
	if err := h.usecase.Submit(ctx, &resign); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Pengajuan resign berhasil"})
}

func (h *ResignationHandler) Approve(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	ctx := c.UserContext()
	if err := h.usecase.Approve(ctx, uint(id)); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Resign disetujui"})
}

func (h *ResignationHandler) GetHistory(c *fiber.Ctx) error {
	ctx := c.UserContext()
	resigns, err := h.usecase.GetHistory(ctx)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resigns)
}
