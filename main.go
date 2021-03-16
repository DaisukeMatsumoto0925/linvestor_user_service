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

type User struct {
	UserName string `json:"user_name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

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
		v1.GET("/user/:id", show)
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
	var userParams User
	err := c.BindJSON(&userParams)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	client := c.MustGet("firebaseAuth").(*auth.Client)

	user := createUser(c, client, userParams)
	c.JSON(http.StatusOK, user)
}

func show(c *gin.Context) {
	client := c.MustGet("firebaseAuth").(*auth.Client)
	id := c.Param("id")
	user := getUser(c, client, id)
	c.JSON(http.StatusOK, user)
}

func createUser(ctx *gin.Context, client *auth.Client, userParams User) *auth.UserRecord {
	params := (&auth.UserToCreate{}).
		Email(userParams.Email).
		Password(userParams.Password).
		DisplayName(userParams.UserName)
	u, err := client.CreateUser(ctx, params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	log.Printf("Successfully created user: %#v\n", u.UserInfo)

	return u
}

func getUser(ctx *gin.Context, client *auth.Client, uid string) *auth.UserRecord {
	u, err := client.GetUser(ctx, uid)
	if err != nil {
		log.Fatalf("error getting user %s: %v\n", uid, err)
	}
	log.Printf("Successfully fetched user data: %v\n", u)

	return u
}
