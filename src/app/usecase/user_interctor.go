package usecase

import (
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/xfpng345/linvestor_user_service/src/app/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(ctx *gin.Context, u domain.User) (user *auth.UserRecord, err error) {
	user, err = interactor.UserRepository.Store(ctx, u)
	return
}
