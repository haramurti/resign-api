package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type LeaveRequest struct {
	gorm.Model
	UserID    uint      `gorm:"not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"-"` // Relation ke User
	Reason    string    `gorm:"type:text;not null" json:"reason"`
	StartDate time.Time `gorm:"not null" json:"start_date"`
	EndDate   time.Time `gorm:"not null" json:"end_date"`
	Status    string    `gorm:"type:varchar(20);default:'pending'" json:"status"` // pending, approved, rejected
}

type LeaveRepository interface {
	Create(ctx context.Context, leave *LeaveRequest) error
	FetchAll(ctx context.Context) ([]LeaveRequest, error)
	GetByID(ctx context.Context, id uint) (LeaveRequest, error)
	UpdateStatus(ctx context.Context, id uint, status string) error
}

type LeaveUsecase interface {
	Apply(ctx context.Context, leave *LeaveRequest) error // Isinya validasi jatah cuti, dll
	GetHistory(ctx context.Context) ([]LeaveRequest, error)
	ApproveLeave(ctx context.Context, id uint) error
	RejectLeave(ctx context.Context, id uint) error
}
