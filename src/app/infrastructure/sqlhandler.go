package infrastructure

import (
	"context"
	"log"
	"os"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/xfpng345/linvestor_user_service/src/app/domain"
	"github.com/xfpng345/linvestor_user_service/src/app/interfaces/database"
	"google.golang.org/api/option"
)

// SQLHandler is a handler
type SQLHandler struct {
	Conn *auth.Client
}

func NewSQLHandler() database.SQLHandler {
	serviceAccountKeyFilePath, err := filepath.Abs(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	if err != nil {
		log.Print("Unable to load serviceAccountKeys.json file")
	}
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Print("Firebase load error")
	}
	conn, err := app.Auth(context.Background())
	if err != nil {
		log.Print("Firebase load error")
	}

	SQLHandler := new(SQLHandler)
	SQLHandler.Conn = conn
	return SQLHandler
}

func (handler *SQLHandler) PostUser(ctx context.Context, u domain.User) (user *auth.UserRecord, err error) {
	params := (&auth.UserToCreate{}).
		Email(u.Email).
		Password(u.Password).
		DisplayName(u.UserName)
	user, err = handler.Conn.CreateUser(ctx, params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	return
}

func (handler *SQLHandler) GetUser(ctx context.Context, uid string) (user *auth.UserRecord, err error) {
	user, err = handler.Conn.GetUser(ctx, uid)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	return
}

func (handler *SQLHandler) DeleteUser(ctx context.Context, uid string) (err error) {
	err = handler.Conn.DeleteUser(ctx, uid)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	return
}

func (handler *SQLHandler) PutUser(ctx context.Context, uid string, u domain.User) (user *auth.UserRecord, err error) {
	params := (&auth.UserToUpdate{}).
		Email(u.Email).
		Password(u.Password).
		DisplayName(u.UserName)
	user, err = handler.Conn.UpdateUser(ctx, uid, params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	return
}
