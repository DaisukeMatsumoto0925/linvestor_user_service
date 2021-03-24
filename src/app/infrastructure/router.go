package infrastructure

import (
	"context"
	"log"
	"os"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	gin "github.com/gin-gonic/gin"
	"github.com/xfpng345/linvestor_user_service/src/app/interfaces/controllers"
	"google.golang.org/api/option"
)

var Router *gin.Engine

func init() {
	engine := gin.Default()
	firebaseAuth := setupFirebase()
	engine.Use(func(c *gin.Context) {
		c.Set("firebaseAuth", firebaseAuth)
	})
	v1 := engine.Group("/v1")
	{
		v1.POST("/users", controllers.Create)
		v1.GET("/user/:id", controllers.Show)
		v1.DELETE("/user/:id", controllers.Delete)
	}
	log.Fatal(engine.Run(":8080"))
}

func setupFirebase() *auth.Client {
	serviceAccountKeyFilePath, err := filepath.Abs(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	if err != nil {
		log.Print("Unable to load serviceAccountKeys.json file")
	}
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Print("Firebase load error")
	}
	auth, err := app.Auth(context.Background())
	if err != nil {
		log.Print("Firebase load error")
	}
	return auth
}
