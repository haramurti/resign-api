package usecase

import (
	"context"
	"errors"
	"resign-api/internal/domain"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(ur domain.UserRepository) domain.UserUsecase {
	return &userUsecase{userRepo: ur}
}

func (u *userUsecase) Register(ctx context.Context, user *domain.User) error {
	// 1. Cek apakah email sudah terdaftar
	existingUser, _ := u.userRepo.GetByEmail(ctx, user.Email)
	if existingUser.ID != 0 {
		return errors.New("email sudah digunakan, pakai email lain bos")
	}

	// 2. Set default quota kalau belum diisi
	if user.LeaveQuota == 0 {
		user.LeaveQuota = 12
	}

	return u.userRepo.Create(ctx, user)
}

func (u *userUsecase) GetProfile(ctx context.Context, id uint) (domain.User, error) {
	return u.userRepo.GetByID(ctx, id)
}

func (u *userUsecase) UpdateQuota(ctx context.Context, id uint, newQuota int) error {
	user, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	user.LeaveQuota = newQuota
	return u.userRepo.Update(ctx, &user)
}
