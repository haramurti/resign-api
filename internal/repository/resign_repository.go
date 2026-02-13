package repository

import (
	"context"
	"resign-api/internal/domain"

	"gorm.io/gorm"
)

type resignationRepository struct {
	db *gorm.DB
}

func NewResignationRepository(db *gorm.DB) domain.ResignationRepository {
	return &resignationRepository{db}
}

func (r *resignationRepository) Create(ctx context.Context, resign *domain.Resignation) error {
	return r.db.WithContext(ctx).Create(resign).Error
}

func (r *resignationRepository) FetchAll(ctx context.Context) ([]domain.Resignation, error) {
	var resigns []domain.Resignation
	err := r.db.WithContext(ctx).Preload("User").Find(&resigns).Error
	return resigns, err
}

func (r *resignationRepository) GetByID(ctx context.Context, id uint) (domain.Resignation, error) {
	var resign domain.Resignation
	err := r.db.WithContext(ctx).Preload("User").First(&resign, id).Error
	return resign, err
}

func (r *resignationRepository) UpdateStatus(ctx context.Context, id uint, status string) error {
	return r.db.WithContext(ctx).Model(&domain.Resignation{}).Where("id = ?", id).Update("status", status).Error
}
