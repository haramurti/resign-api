package domain

import (
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
