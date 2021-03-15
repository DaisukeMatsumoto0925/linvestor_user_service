package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func main() {
	engine := gin.Default()
	firebaseAuth := setupFirebase()
	engine.Use(func(c *gin.Context) {
		c.Set("firebaseAuth", firebaseAuth)
	})
	userEngine := engine.Group("/user")
	{
		v1 := userEngine.Group("/v1")
		{
			v1.GET("/", public)
		}
	}
	log.Fatal(engine.Run(":8080"))
}

func public(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
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
