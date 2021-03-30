package usecase

import (
	"context"

	"github.com/xfpng345/linvestor_user_service/src/app/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, u domain.User) (domain.User, error)
	GetByID(ctx context.Context, u string) (domain.User, error)
	DeleteUser(ctx context.Context, u string) error
	UpdateUser(ctx context.Context, uid string, u domain.User) (domain.User, error)
}
