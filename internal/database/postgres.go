package database

import (
	"log"
	"os"
	"resign-api/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set!")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: false, // MATIKAN INI
	})

	if err != nil {
		log.Fatalf("Failed to connect to supabase via GORM: %v", err)
	}

	log.Println("GORM connected successfully to supabase!")

	log.Println("Running Auto Migration...")
	err = db.AutoMigrate(
		&domain.User{},
		&domain.LeaveRequest{},
		&domain.Resignation{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("✅ Database schema is up to date!")

	return db
}
