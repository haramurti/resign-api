package handler

import (
	"encoding/base64"
	"resign-api/internal/domain"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Method 1: Cek Login Gmail & Password dari Database
func NewAuthMiddleware(userRepo domain.UserRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Basic ") {
			return c.Status(401).JSON(fiber.Map{"error": "Login required"})
		}

		// Decode Basic Auth manual
		payload, _ := base64.StdEncoding.DecodeString(auth[6:])
		pair := strings.SplitN(string(payload), ":", 2)
		if len(pair) != 2 {
			return c.Status(401).JSON(fiber.Map{"error": "Format auth salah"})
		}

		email, password := pair[0], pair[1]
		ctx := c.UserContext()
		user, err := userRepo.GetByEmail(ctx, email)

		// Validasi dengan data di struct User Supabase
		if err != nil || user.Password != password {
			return c.Status(401).JSON(fiber.Map{"error": "Email/Password salah!"})
		}

		// Simpen user ke locals buat dipake method AdminOnly
		c.Locals("currentUser", user)
		return c.Next()
	}
}

// Method 2: Cek Role Manager atau HR
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
