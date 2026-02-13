package usecase

import (
	"context"
	"errors"
	"resign-api/internal/domain"
)

type resignationUsecase struct {
	resignRepo domain.ResignationRepository
}

func NewResignationUsecase(rr domain.ResignationRepository) domain.ResignationUsecase {
	return &resignationUsecase{resignRepo: rr}
}

func (u *resignationUsecase) Submit(ctx context.Context, resign *domain.Resignation) error {
	resign.Status = "pending"
	return u.resignRepo.Create(ctx, resign)
}

func (u *resignationUsecase) GetHistory(ctx context.Context) ([]domain.Resignation, error) {
	return u.resignRepo.FetchAll(ctx)
}

func (u *resignationUsecase) Approve(ctx context.Context, id uint) error {
	resign, err := u.resignRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if resign.Status == "approved" {
		return errors.New("karyawan ini sudah resmi resign, nggak bisa approve dua kali")
	}

	return u.resignRepo.UpdateStatus(ctx, id, "approved")
}

func (u *resignationUsecase) Reject(ctx context.Context, id uint) error {
	return u.resignRepo.UpdateStatus(ctx, id, "rejected")
}
