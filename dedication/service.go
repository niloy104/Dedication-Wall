package dedication

import (
	"context"
	"dedicationWall/domain"
)

type service struct {
	dedRepo DedicationRepo
}

func NewService(dedRepo DedicationRepo) Service {
	return &service{
		dedRepo: dedRepo,
	}
}

func (s *service) Create(ctx context.Context, d domain.Dedication) (*domain.Dedication, error) {
	return s.dedRepo.Create(ctx, d)
}

func (s *service) List(ctx context.Context, page, limit int64) ([]*domain.Dedication, int64, error) {
	return s.dedRepo.List(ctx, page, limit)
}

func (s *service) GetByID(ctx context.Context, id uint64) (*domain.Dedication, error) {
	return s.dedRepo.GetByID(ctx, id)
}


func (s *service) Delete(ctx context.Context, id uint64) error {
	return s.dedRepo.Delete(ctx, id)
}