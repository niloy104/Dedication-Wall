// user/port.go
package user

import (
	"context"
	"dedicationWall/domain"
	userHandler "dedicationWall/rest/handlers/user"
)

type Service interface {
	userHandler.Service // embedding — exactly like your old project
}

type UserRepo interface {
	Create(ctx context.Context, user domain.User) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
}