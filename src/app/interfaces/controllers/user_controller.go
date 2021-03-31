package controllers

import (
	"net/http"

	"github.com/xfpng345/linvestor_user_service/src/app/domain"
	"github.com/xfpng345/linvestor_user_service/src/app/interfaces/database"
	"github.com/xfpng345/linvestor_user_service/src/app/usecase"
)

type jsonResponse struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

func responseJSON(responseCode int, responseMsg string) jsonResponse {
	return jsonResponse{
		Status:  responseCode,
		Message: responseMsg,
	}
}

// UserController is struct of controller
type UserController struct {
	Interactor usecase.UserInteractor
}

// NewUserController is constructor
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
		c.JSON(http.StatusBadRequest, responseJSON(http.StatusBadRequest, err.Error()))
		return
	}

	user, err := controller.Interactor.UserCreate(c, u)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseJSON(http.StatusBadRequest, err.Error()))
	}
	c.JSON(http.StatusOK, user)
}

// Show func get a user
func (controller *UserController) Show(c Context) {
	id := c.Param("id")
	user, err := controller.Interactor.UserByID(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseJSON(http.StatusBadRequest, err.Error()))
	}
	c.JSON(http.StatusOK, user)
}

// Delete func delete a user
func (controller *UserController) Delete(c Context) {
	id := c.Param("id")
	err := controller.Interactor.DeleteByID(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseJSON(http.StatusBadRequest, err.Error()))
	}
	c.JSON(http.StatusOK, responseJSON(http.StatusOK, "Deleted a user"))
}

// Update func update a user
func (controller *UserController) Update(c Context) {
	id := c.Param("id")
	var u domain.User
	err := c.BindJSON(&u)
	user, err := controller.Interactor.UpdateByID(c, id, u)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseJSON(http.StatusBadRequest, err.Error()))
	}
	c.JSON(http.StatusOK, user)
}
