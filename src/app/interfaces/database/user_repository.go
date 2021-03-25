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
