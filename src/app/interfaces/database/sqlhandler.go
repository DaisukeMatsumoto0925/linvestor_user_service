package database

import (
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/xfpng345/linvestor_user_service/src/app/domain"
)

type SqlHandler interface {
	PostUser(ctx *gin.Context, u domain.User) (user *auth.UserRecord, err error)
	GetUser(ctx *gin.Context, uid string) (user *auth.UserRecord, err error)
	DeleteUser(ctx *gin.Context, uid string) error
}
