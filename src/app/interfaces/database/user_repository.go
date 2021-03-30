package database

import (
	"context"
	"log"

	"firebase.google.com/go/auth"
	"github.com/xfpng345/linvestor_user_service/src/app/domain"
)

// UserRepository is interface
type UserRepository struct {
	SQLHandler
}

// Store is func save a user
func (repo *UserRepository) Store(ctx context.Context, u domain.User) (*auth.UserRecord, error) {
	user, err := repo.PostUser(ctx, u)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	log.Printf("Successfully created user: %#v\n", user.UserInfo)

	return user, nil
}

// GetByID is func get a user
func (repo *UserRepository) GetByID(ctx context.Context, uid string) (user *auth.UserRecord, err error) {
	user, err = repo.GetUser(ctx, uid)
	if err != nil {
		log.Fatalf("error getting user %s: %v\n", uid, err)
	}
	log.Printf("Successfully fetched user data: %v\n", user)

	return user, nil
}

// DeleteByID is func delete a user
func (repo *UserRepository) DeleteByID(ctx context.Context, uid string) (err error) {
	err = repo.DeleteUser(ctx, uid)
	if err != nil {
		return err
	}
	log.Printf("Successfully deleted user: %s\n", uid)
	return nil
}

// UpdateUser is func update a user
func (repo *UserRepository) UpdateUser(ctx context.Context, uid string, u domain.User) (*auth.UserRecord, error) {
	user, err := repo.PutUser(ctx, uid, u)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	log.Printf("Successfully created user: %#v\n", user.UserInfo)

	return user, nil
}
