package usecase

import (
	"context"
	"errors"
	"resign-api/internal/domain"
)

type leaveUsecase struct {
	leaveRepo domain.LeaveRepository
	userRepo  domain.UserRepository
}

func NewLeaveUsecase(lr domain.LeaveRepository, ur domain.UserRepository) domain.LeaveUsecase {
	return &leaveUsecase{
		leaveRepo: lr,
		userRepo:  ur,
	}
}

func (u *leaveUsecase) Apply(ctx context.Context, leave *domain.LeaveRequest) error {
	user, err := u.userRepo.GetByID(ctx, leave.UserID)
	if err != nil {
		return errors.New("user tidak ditemukan")
	}

	if user.LeaveQuota <= 0 {
		return errors.New("jatah cuti sudah habis, kerja terus sampe tipes!")
	}

	leave.Status = "pending"
	return u.leaveRepo.Create(ctx, leave)
}

func (u *leaveUsecase) GetHistory(ctx context.Context) ([]domain.LeaveRequest, error) {
	return u.leaveRepo.FetchAll(ctx)
}

func (u *leaveUsecase) ApproveLeave(ctx context.Context, id uint) error {
	leave, err := u.leaveRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if leave.Status != "pending" {
		return errors.New("hanya bisa approve pengajuan yang statusnya pending")
	}

	err = u.leaveRepo.UpdateStatus(ctx, id, "approved")
	if err != nil {
		return err
	}

	user, _ := u.userRepo.GetByID(ctx, leave.UserID)
	user.LeaveQuota -= 1
	return u.userRepo.Update(ctx, &user)
}

func (u *leaveUsecase) RejectLeave(ctx context.Context, id uint) error {
	return u.leaveRepo.UpdateStatus(ctx, id, "rejected")
}
