package database

import (
	"context"
	"log"

	"github.com/xfpng345/linvestor_user_service/src/app/domain"
)

// UserRepository is interface
type UserRepository struct {
	SQLHandler
}

// CreateUser is func save a user
func (repo *UserRepository) CreateUser(ctx context.Context, u domain.User) (user domain.User, err error) {
	user, err = repo.PostUser(ctx, u)
	if err != nil {
		return
	}
	log.Printf("Successfully created user: %#v\n", user)

	return
}

// GetByID is func get a user
func (repo *UserRepository) GetByID(ctx context.Context, uid string) (user domain.User, err error) {
	user, err = repo.GetUser(ctx, uid)
	if err != nil {
		return
	}
	log.Printf("Successfully fetched user data: %v\n", user)

	return
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
func (repo *UserRepository) UpdateUser(ctx context.Context, uid string, u domain.User) (user domain.User, err error) {
	user, err = repo.PutUser(ctx, uid, u)
	if err != nil {
		return
	}
	log.Printf("Successfully created user: %#v\n", user)

	return
}
