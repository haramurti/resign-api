package domain

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string `gorm:"type:varchar(100);not null" json:"name"`
	Email      string `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Role       string `gorm:"type:varchar(20);default:'employee'" json:"role"` // employee, manager, hr
	LeaveQuota int    `gorm:"default:12" json:"leave_quota"`                   // Jatah cuti tahunan
}

// UserRepository: Kontrak buat CRUD User di database
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id uint) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id uint) error
}

// UserUsecase: Kontrak buat logic User
type UserUsecase interface {
	Register(ctx context.Context, user *User) error
	GetProfile(ctx context.Context, id uint) (User, error)
	UpdateQuota(ctx context.Context, id uint, newQuota int) error
}
