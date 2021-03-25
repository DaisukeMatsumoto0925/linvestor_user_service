package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xfpng345/linvestor_user_service/src/app/domain"
	"github.com/xfpng345/linvestor_user_service/src/app/interfaces/database"
	"github.com/xfpng345/linvestor_user_service/src/app/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

// Create func post a user
func (controller *UserController) Create(c *gin.Context) {
	var userParams domain.User
	err := c.BindJSON(&userParams)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	user, _ := controller.Interactor.Add(c, userParams)
	c.JSON(http.StatusOK, user)
}

// Show func get a user
func (controller *UserController) Show(c *gin.Context) {
	id := c.Param("id")
	user, _ := controller.Interactor.UserByID(c, id)
	c.JSON(http.StatusOK, user)
}

// // Delete func delete a user
// func Delete(c *gin.Context) {
// 	client := c.MustGet("firebaseAuth").(*auth.Client)
// 	id := c.Param("id")
// 	err := deleteUser(c, client, id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err.Error())
// 	}
// 	c.JSON(http.StatusOK, "ok")
// }

// func getUser(ctx *gin.Context, client *auth.Client, uid string) *auth.UserRecord {
// 	u, err := client.GetUser(ctx, uid)
// 	if err != nil {
// 		log.Fatalf("error getting user %s: %v\n", uid, err)
// 	}
// 	log.Printf("Successfully fetched user data: %v\n", u)

// 	return u
// }

// func deleteUser(ctx *gin.Context, client *auth.Client, uid string) error {
// 	err := client.DeleteUser(ctx, uid)
// 	if err != nil {
// 		return err
// 	}
// 	log.Printf("Successfully deleted user: %s\n", uid)
// 	return nil
// }
