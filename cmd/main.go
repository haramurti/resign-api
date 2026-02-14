package main

import (
	"log"
	"os"
	"resign-api/internal/database"
	"resign-api/internal/handler"
	"resign-api/internal/repository"
	"resign-api/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	// 1. Koneksi Database (Sekaligus AutoMigrate di dalamnya)
	db := database.NewPostgresDB()

	// 2. Wiring Up (Menyambungkan Kabel)
	// --- USER ---
	userRepo := repository.NewUserRepository(db)
	userUC := usecase.NewUserUsecase(userRepo)
	userHdl := handler.NewUserHandler(userUC)

	// --- LEAVE ---
	leaveRepo := repository.NewLeaveRepository(db)
	leaveUC := usecase.NewLeaveUsecase(leaveRepo, userRepo) // Butuh userRepo buat cek kuota
	leaveHdl := handler.NewLeaveHandler(leaveUC)

	// --- RESIGN ---
	resignRepo := repository.NewResignationRepository(db)
	resignUC := usecase.NewResignationUsecase(resignRepo)
	resignHdl := handler.NewResignationHandler(resignUC)

	app := fiber.New()

	app.Static("/", "./public")

	// 3. Routing
	api := app.Group("/api") // Grouping biar rapi /api/...

	// Route User
	api.Post("/users", userHdl.Register)
	api.Get("/users/:id", userHdl.GetProfile)

	// Route Leave
	api.Post("/leaves", leaveHdl.Apply)
	api.Get("/leaves", leaveHdl.GetHistory)
	api.Patch("/leaves/:id/approve", leaveHdl.Approve)

	// Route Resign
	api.Post("/resignations", resignHdl.Submit)
	api.Patch("/resignations/:id/approve", resignHdl.Approve)

	// Health Check
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "OK", "message": "Backend is ready to rock!"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8686" // Default port kalau di .env gak ada
	}

	log.Printf("🚀 Server is starting on port %s...", port)
	log.Fatal(app.Listen(":" + port))
}
