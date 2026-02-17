package handler

import (
	"encoding/base64"
	"resign-api/internal/domain"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func NewAuthMiddleware(userRepo domain.UserRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusNoContent)
		}

		auth := c.Get("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Basic ") {
			return c.Status(401).JSON(fiber.Map{"error": "Login required"})
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[6:])
		pair := strings.SplitN(string(payload), ":", 2)
		if len(pair) != 2 {
			return c.Status(401).JSON(fiber.Map{"error": "Format auth salah"})
		}

		email, password := pair[0], pair[1]
		ctx := c.UserContext()
		user, err := userRepo.GetByEmail(ctx, email)

		if err != nil || user.Password != password {
			return c.Status(401).JSON(fiber.Map{"error": "Email/Password salah!"})
		}

		c.Locals("currentUser", user)
		return c.Next()
	}
}

// checking 'hr' role for admin access
func AdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRaw := c.Locals("currentUser")
		if userRaw == nil {
			return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
		}

		user := userRaw.(domain.User)
		if user.Role == "manager" || user.Role == "hr" {
			return c.Next()
		}
		return c.Status(403).JSON(fiber.Map{"error": "Hanya Manager/HR yang bisa akses!"})
	}
}
