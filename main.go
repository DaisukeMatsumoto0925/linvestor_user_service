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
	v1 := engine.Group("/v1")
	{
		v1.GET("/", public)
		v1.POST("/users", create)
		v1.GET("/user", show)
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

func create(c *gin.Context) {
	client := setupFirebase()
	user := createUser(c, client)
	c.JSON(http.StatusOK, user)
}

func show(c *gin.Context) {
	client := setupFirebase()
	email := "user@example.com"
	user := getUserByEmail(c, client, email)
	c.JSON(http.StatusOK, user)
}

func createUser(ctx *gin.Context, client *auth.Client) *auth.UserRecord {
	params := (&auth.UserToCreate{}).
		Email("user@example.com").
		EmailVerified(false).
		PhoneNumber("+15555550100").
		Password("secretPassword").
		DisplayName("John Doe").
		PhotoURL("http://www.example.com/12345678/photo.png").
		Disabled(false)
	u, err := client.CreateUser(ctx, params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	log.Printf("Successfully created user: %#v\n", u.UserInfo)

	return u
}

func getUserByEmail(ctx *gin.Context, client *auth.Client, email string) *auth.UserRecord {
	u, err := client.GetUserByEmail(ctx, email)
	if err != nil {
		log.Fatalf("error getting user by email %s: %v\n", email, err)
	}
	log.Printf("Successfully fetched user data: %v\n", u)
	return u
}
