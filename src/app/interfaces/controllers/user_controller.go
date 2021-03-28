package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xfpng345/linvestor_user_service/src/app/domain"
	"github.com/xfpng345/linvestor_user_service/src/app/interfaces/database"
	"github.com/xfpng345/linvestor_user_service/src/app/usecase"
)

// UserController is struct of controller
type UserController struct {
	Interactor usecase.UserInteractor
}

// NewUserController si constructor
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

// Delete func delete a user
func (controller *UserController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := controller.Interactor.DeleteByID(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, "ok")
}
