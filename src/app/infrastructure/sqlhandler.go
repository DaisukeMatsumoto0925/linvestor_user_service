package infrastructure

import (
	"context"
	"log"
	"os"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	gin "github.com/gin-gonic/gin"
	"github.com/xfpng345/linvestor_user_service/src/app/domain"
	"github.com/xfpng345/linvestor_user_service/src/app/interfaces/database"
	"google.golang.org/api/option"
)

// SqlHandler is a handler
type SqlHandler struct {
	Conn *auth.Client
}

func NewSqlHandler() database.SqlHandler {
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

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func (handler *SqlHandler) PostUser(ctx *gin.Context, u domain.User) (user *auth.UserRecord, err error) {
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
