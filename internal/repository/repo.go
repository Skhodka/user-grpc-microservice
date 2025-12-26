package repository

import (
	"context"
	domain "usermic/internal/domain/user"
)

type UserRepo interface {
	RegUser(ctx context.Context, us *domain.UserDomain) (int, error)
	FindByEmail(ctx context.Context, email string) (*domain.UserDomain, error)
	FindByPhone(ctx context.Context, phone string) (*domain.UserDomain, error)
}
