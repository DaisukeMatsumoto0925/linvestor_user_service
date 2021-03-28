package database

import (
	"log"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/xfpng345/linvestor_user_service/src/app/domain"
)

type UserRepository struct {
	SqlHandler
}

// Store is func save a user
func (repo *UserRepository) Store(ctx *gin.Context, u domain.User) (*auth.UserRecord, error) {
	user, err := repo.PostUser(ctx, u)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	log.Printf("Successfully created user: %#v\n", user.UserInfo)

	return user, nil
}

func (repo *UserRepository) GetByID(ctx *gin.Context, uid string) (user *auth.UserRecord, err error) {
	user, err = repo.GetUser(ctx, uid)
	if err != nil {
		log.Fatalf("error getting user %s: %v\n", uid, err)
	}
	log.Printf("Successfully fetched user data: %v\n", user)

	return user, nil
}

func (repo *UserRepository) DeleteByID(ctx *gin.Context, uid string) (err error) {
	err = repo.DeleteUser(ctx, uid)
	if err != nil {
		return err
	}
	log.Printf("Successfully deleted user: %s\n", uid)
	return nil
}

func (repo *UserRepository) UpdateUser(ctx *gin.Context, uid string, u domain.User) (*auth.UserRecord, error) {
	user, err := repo.PutUser(ctx, uid, u)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	log.Printf("Successfully created user: %#v\n", user.UserInfo)

	return user, nil
}
