package dedication

import (
	"context"
	"dedicationWall/domain"
	handler "dedicationWall/rest/handlers/dedication"
)

type Service interface {
	handler.Service 
}

type DedicationRepo interface {
	Create(ctx context.Context, d domain.Dedication) (*domain.Dedication, error)
	List(ctx context.Context, page, limit int64) ([]*domain.Dedication, int64, error)
	GetByID(ctx context.Context, id uint64) (*domain.Dedication, error)
	Delete(ctx context.Context, id uint64) error
}
