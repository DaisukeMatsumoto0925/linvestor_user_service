package controllers

import (
	"log"
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/xfpng345/linvestor_user_service/src/app/domain"
)

// Create func post a user
func Create(c *gin.Context) {
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

// Show func get a user
func Show(c *gin.Context) {
	client := c.MustGet("firebaseAuth").(*auth.Client)
	id := c.Param("id")
	user := getUser(c, client, id)
	c.JSON(http.StatusOK, user)
}

// Delete func delete a user
func Delete(c *gin.Context) {
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
