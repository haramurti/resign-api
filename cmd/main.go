package main

import (
	"log"
	"os"
	"resign-api/internal/database"
	"resign-api/internal/handler"
	"resign-api/internal/repository"
	"resign-api/internal/usecase"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	db := database.NewPostgresDB()

	//repositories
	userRepo := repository.NewUserRepository(db)
	leaveRepo := repository.NewLeaveRepository(db)
	resignRepo := repository.NewResignationRepository(db)

	//usecases
	userUC := usecase.NewUserUsecase(userRepo)
	leaveUC := usecase.NewLeaveUsecase(leaveRepo, userRepo)
	resignUC := usecase.NewResignationUsecase(resignRepo)

	//hanlders
	userHdl := handler.NewUserHandler(userUC)
	leaveHdl := handler.NewLeaveHandler(leaveUC)
	resignHdl := handler.NewResignationHandler(resignUC)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PATCH, DELETE, OPTIONS",
	}))
	app.Use(logger.New())
	app.Static("/", "./public")

	//middlewares
	authMid := handler.NewAuthMiddleware(userRepo)
	adminMid := handler.AdminOnly()

	// routes
	//need login all
	api := app.Group("/api", authMid)

	// User Routes
	api.Get("/me", userHdl.GetProfile) // Biar userHdl kepake

	// Employee Routes
	api.Post("/leaves", leaveHdl.Apply)
	api.Get("/leaves", leaveHdl.GetHistory)
	api.Post("/resignations", resignHdl.Submit)
	api.Get("/resignations", resignHdl.GetHistory)

	//admin routes 'hr' roles only
	admin := api.Group("/admin", adminMid)
	admin.Patch("/leaves/:id/approve", leaveHdl.Approve)
	admin.Patch("/resignations/:id/approve", resignHdl.Approve)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8686"
	}

	log.Printf("🚀 BCA System started on port %s...", port)
	log.Fatal(app.Listen(":" + port))
}

//apply working

//test

//test

//test 3
