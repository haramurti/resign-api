package domain

import (
	"time"

	"gorm.io/gorm"
)

type Resignation struct {
	gorm.Model
	UserID     uint      `gorm:"not null" json:"user_id"`
	User       User      `gorm:"foreignKey:UserID" json:"-"`
	ResignDate time.Time `gorm:"not null" json:"resign_date"`
	Reason     string    `gorm:"type:text;not null" json:"reason"`
	Status     string    `gorm:"type:varchar(20);default:'pending'" json:"status"`
}
