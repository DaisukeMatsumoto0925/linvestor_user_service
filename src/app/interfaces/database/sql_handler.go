package database

import (
	"context"

	"github.com/xfpng345/linvestor_user_service/src/app/domain"
)

// SQLHandler is a interface of handler
type SQLHandler interface {
	PostUser(ctx context.Context, u domain.User) (user domain.User, err error)
	GetUser(ctx context.Context, uid string) (user domain.User, err error)
	DeleteUser(ctx context.Context, uid string) error
	PutUser(ctx context.Context, uid string, u domain.User) (user domain.User, err error)
}
