package database

import (
	"log"
	"os"
	"resign-api/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresDB() *gorm.DB {
	//getting database url supabase
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set!")
	}

	//open connection on gorm
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Failed to connect to supabase via GORM: %v", err)
	}

	log.Println("GORM connected successfully to supabase!")

	//automigration
	log.Println("Running Auto Migration...")
	err = db.AutoMigrate(
		&domain.User{},
		&domain.LeaveRequest{},
		&domain.Resignation{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database is now up to date!")

	return db
}
