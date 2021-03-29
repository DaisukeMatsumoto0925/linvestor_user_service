package usecase

import (
	"context"

	"firebase.google.com/go/auth"
	"github.com/xfpng345/linvestor_user_service/src/app/domain"
)

type UserRepository interface {
	Store(ctx context.Context, u domain.User) (*auth.UserRecord, error)
	GetByID(ctx context.Context, u string) (*auth.UserRecord, error)
	DeleteUser(ctx context.Context, u string) error
	UpdateUser(ctx context.Context, uid string, u domain.User) (*auth.UserRecord, error)
}
