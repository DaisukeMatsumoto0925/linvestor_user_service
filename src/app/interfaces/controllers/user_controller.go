package controllers

import (
	"net/http"

	"github.com/xfpng345/linvestor_user_service/src/app/domain"
	"github.com/xfpng345/linvestor_user_service/src/app/interfaces/database"
	"github.com/xfpng345/linvestor_user_service/src/app/usecase"
)

// UserController is struct of controller
type UserController struct {
	Interactor usecase.UserInteractor
}

// NewUserController si constructor
func NewUserController(SQLHandler database.SQLHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SQLHandler: SQLHandler,
			},
		},
	}
}

// Create func post a user
func (controller *UserController) Create(c Context) {
	var u domain.User
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	user, _ := controller.Interactor.Add(c, u)
	c.JSON(http.StatusOK, user)
}

// Show func get a user
func (controller *UserController) Show(c Context) {
	id := c.Param("id")
	user, _ := controller.Interactor.UserByID(c, id)
	c.JSON(http.StatusOK, user)
}

// Delete func delete a user
func (controller *UserController) Delete(c Context) {
	id := c.Param("id")
	err := controller.Interactor.DeleteByID(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, "ok")
}

// Update func update a user
func (controller *UserController) Update(c Context) {
	id := c.Param("id")
	var u domain.User
	err := c.BindJSON(&u)
	user, err := controller.Interactor.UpdateByID(c, id, u)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, user)
}
