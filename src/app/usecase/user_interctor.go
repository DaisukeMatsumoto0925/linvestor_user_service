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

func (interactor *UserInteractor) UserByID(ctx *gin.Context, uid string) (user *auth.UserRecord, err error) {
	user, err = interactor.UserRepository.GetByID(ctx, uid)
	return
}

func (interactor *UserInteractor) DeleteByID(ctx *gin.Context, uid string) (err error) {
	err = interactor.UserRepository.DeleteUser(ctx, uid)
	if err != nil {
		return err
	}
	return nil
}
