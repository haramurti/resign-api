package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Resignation struct {
	gorm.Model
	UserID     uint      `gorm:"not null" json:"user_id"`
	User       User      `gorm:"foreignKey:UserID" json:"User"`
	ResignDate time.Time `gorm:"not null" json:"resign_date"`
	Reason     string    `gorm:"type:text;not null" json:"reason"`
	Status     string    `gorm:"type:varchar(20);default:'pending'" json:"status"`
}

type ResignationRepository interface {
	Create(ctx context.Context, resign *Resignation) error
	FetchAll(ctx context.Context) ([]Resignation, error)
	GetByID(ctx context.Context, id uint) (Resignation, error)
	UpdateStatus(ctx context.Context, id uint, status string) error
}

type ResignationUsecase interface {
	Submit(ctx context.Context, resign *Resignation) error
	GetHistory(ctx context.Context) ([]Resignation, error)
	Approve(ctx context.Context, id uint) error
	Reject(ctx context.Context, id uint) error
}
