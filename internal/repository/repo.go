package repository

import (
	"context"
)

type UserRepo interface {
	RegUser(ctx context.Context, us *UserRecord) (int, error)
	FindByEmail(ctx context.Context, email string) (*UserRecord, error)
	FindByPhone(ctx context.Context, phone string) (*UserRecord, error)
}
