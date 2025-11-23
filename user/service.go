package user

import (
	"context"
	"dedicationWall/domain"
)

type service struct {
	usrRepo UserRepo
}

func NewService(usrRepo UserRepo) Service {
	return &service{
		usrRepo: usrRepo,
	}
}


func (svc *service) Create(ctx context.Context, user domain.User) (*domain.User, error) {
	usr, err := svc.usrRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, nil
	}

	return usr, nil
}

func (svc *service) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	usr, err := svc.usrRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, nil
	}

	return usr, nil
}
