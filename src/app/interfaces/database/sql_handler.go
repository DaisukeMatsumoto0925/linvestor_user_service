package database

import (
	"context"

	"firebase.google.com/go/auth"
	"github.com/xfpng345/linvestor_user_service/src/app/domain"
)

type SqlHandler interface {
	PostUser(ctx context.Context, u domain.User) (user *auth.UserRecord, err error)
	GetUser(ctx context.Context, uid string) (user *auth.UserRecord, err error)
	DeleteUser(ctx context.Context, uid string) error
	PutUser(ctx context.Context, uid string, u domain.User) (user *auth.UserRecord, err error)
}
