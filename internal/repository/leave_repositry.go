package repository

import (
	"context"
	"resign-api/internal/domain"

	"gorm.io/gorm"
)

type leaveRepository struct {
	db *gorm.DB
}

func NewLeaveRepository(db *gorm.DB) domain.LeaveRepository {
	return &leaveRepository{db}
}

func (r *leaveRepository) Create(ctx context.Context, leave *domain.LeaveRequest) error {
	// GORM bakal otomatis ngisi UserID ke tabel leave_requests
	return r.db.WithContext(ctx).Create(leave).Error
}

func (r *leaveRepository) FetchAll(ctx context.Context) ([]domain.LeaveRequest, error) {
	var leaves []domain.LeaveRequest
	// .Preload("User") itu kayak "JOIN" di SQL, biar data usernya ikut ketarik
	err := r.db.WithContext(ctx).Preload("User").Find(&leaves).Error
	return leaves, err
}

func (r *leaveRepository) GetByID(ctx context.Context, id uint) (domain.LeaveRequest, error) {
	var leave domain.LeaveRequest
	err := r.db.WithContext(ctx).Preload("User").First(&leave, id).Error
	return leave, err
}

func (r *leaveRepository) UpdateStatus(ctx context.Context, id uint, status string) error {
	// Update spesifik kolom status aja (buat HR approval)
	return r.db.WithContext(ctx).Model(&domain.LeaveRequest{}).Where("id = ?", id).Update("status", status).Error
}
