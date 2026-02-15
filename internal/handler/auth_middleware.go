package handler

import (
	"encoding/base64"
	"resign-api/internal/domain"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// NewAuthMiddleware: Cek Gmail & Pass dari Database
func NewAuthMiddleware(userRepo domain.UserRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Basic ") {
			return c.Status(401).JSON(fiber.Map{"error": "Login diperlukan"})
		}

		// Decode Basic Auth manual karena Fiber tidak punya c.BasicAuth()
		payload, err := base64.StdEncoding.DecodeString(auth[6:])
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "Format auth salah"})
		}
		pair := strings.SplitN(string(payload), ":", 2)
		if len(pair) != 2 {
			return c.Status(401).JSON(fiber.Map{"error": "Format auth salah"})
		}

		email, password := pair[0], pair[1]
		ctx := c.UserContext()
		user, err := userRepo.GetByEmail(ctx, email)

		// Cocokkan email dan password dari struct User
		if err != nil || user.Password != password {
			return c.Status(401).JSON(fiber.Map{"error": "Credential BCA salah!"})
		}

		// Simpen user ke locals biar bisa dipake middleware AdminOnly
		c.Locals("currentUser", user)
		return c.Next()
	}
}

// AdminOnly: Gatekeeper khusus Role Manager atau HR
func AdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Ambil data user yang udah disimpen tadi di locals
		userRaw := c.Locals("currentUser")
		if userRaw == nil {
			return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
		}

		user := userRaw.(domain.User)

		// Cek Role: Cuma manager atau hr yang boleh lewat
		if user.Role == "manager" || user.Role == "hr" {
			return c.Next()
		}
		return c.Status(403).JSON(fiber.Map{"error": "Terlarang: Anda bukan Manager/HR!"})
	}
}
