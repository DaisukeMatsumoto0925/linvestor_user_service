package usecase

import (
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/xfpng345/linvestor_user_service/src/app/domain"
)

type UserRepository interface {
	Store(ctx *gin.Context, u domain.User) (*auth.UserRecord, error)
}
