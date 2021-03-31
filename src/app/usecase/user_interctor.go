package usecase

import (
	"context"

	"github.com/xfpng345/linvestor_user_service/src/app/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) UserCreate(ctx context.Context, u domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.CreateUser(ctx, u)
	if err != nil {
		return
	}
	return
}

func (interactor *UserInteractor) UserByID(ctx context.Context, uid string) (user domain.User, err error) {
	user, err = interactor.UserRepository.GetByID(ctx, uid)
	if err != nil {
		return
	}
	return
}

func (interactor *UserInteractor) DeleteByID(ctx context.Context, uid string) (err error) {
	err = interactor.UserRepository.DeleteUser(ctx, uid)
	if err != nil {
		return err
	}
	return nil
}

func (interactor *UserInteractor) UpdateByID(ctx context.Context, uid string, u domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.UpdateUser(ctx, uid, u)
	if err != nil {
		return
	}
	return
}
