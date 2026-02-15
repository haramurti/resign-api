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

	// 1. Koneksi Database (AutoMigrate jalan di sini)
	db := database.NewPostgresDB()

	// 2. Wiring Up (Menyambungkan Kabel)

	// --- USER ---
	userRepo := repository.NewUserRepository(db)
	userUC := usecase.NewUserUsecase(userRepo)
	userHdl := handler.NewUserHandler(userUC) // Sekarang bakal kita pake!

	// --- LEAVE ---
	leaveRepo := repository.NewLeaveRepository(db)
	leaveUC := usecase.NewLeaveUsecase(leaveRepo, userRepo)
	leaveHdl := handler.NewLeaveHandler(leaveUC)

	// --- RESIGN ---
	resignRepo := repository.NewResignationRepository(db)
	resignUC := usecase.NewResignationUsecase(resignRepo)
	resignHdl := handler.NewResignationHandler(resignUC)

	app := fiber.New()

	// Serve Static Files (Frontend lo di sini)
	app.Static("/", "./public")

	// 3. Middlewares
	authMid := handler.NewAuthMiddleware(userRepo) // Cek login email & pass
	adminMid := handler.AdminOnly()                // Cek role manager/hr

	// 4. Routing

	// Group API: Semua rute di bawah ini wajib login pake email & password
	api := app.Group("/api", authMid)

	// --- USER ROUTES ---
	api.Get("/profile/:id", userHdl.GetProfile) // userHdl dipake di sini!

	// --- EMPLOYEE ROUTES (Self Service) ---
	api.Post("/leaves", leaveHdl.Apply)
	api.Get("/leaves", leaveHdl.GetHistory)
	api.Post("/resignations", resignHdl.Submit)
	api.Get("/resignations", resignHdl.GetHistory)

	// --- ADMIN ROUTES (MANAGER/HR ONLY) ---
	// Dibungkus lagi pake adminMid biar cuma bos yang bisa masuk
	admin := api.Group("/admin", adminMid)
	admin.Patch("/leaves/:id/approve", leaveHdl.Approve)
	admin.Patch("/resignations/:id/approve", resignHdl.Approve)

	// Health Check buat mastiin backend idup
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "OK", "message": "BCA Backend is Ready to Rock!"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8686"
	}

	log.Printf("🚀 Server is starting on port %s...", port)
	log.Fatal(app.Listen(":" + port))
}
