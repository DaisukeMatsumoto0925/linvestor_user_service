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
	"github.com/xfpng345/linvestor_user_service/domain"
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
		v1.POST("/users", create)
		v1.GET("/user/:id", show)
		v1.DELETE("/user/:id", delete)
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

func create(c *gin.Context) {
	var userParams domain.User
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

func delete(c *gin.Context) {
	client := c.MustGet("firebaseAuth").(*auth.Client)
	id := c.Param("id")
	err := deleteUser(c, client, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, "ok")
}

func createUser(ctx *gin.Context, client *auth.Client, userParams domain.User) *auth.UserRecord {
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

func deleteUser(ctx *gin.Context, client *auth.Client, uid string) error {
	err := client.DeleteUser(ctx, uid)
	if err != nil {
		return err
	}
	log.Printf("Successfully deleted user: %s\n", uid)
	return nil
}
