package infrastructure

import (
	"log"

	gin "github.com/gin-gonic/gin"
	"github.com/xfpng345/linvestor_user_service/src/app/interfaces/controllers"
)

// Router is gin engine
var Router *gin.Engine

func init() {
	engine := gin.Default()
	v1 := engine.Group("/v1")
	userController := controllers.NewUserController(NewSQLHandler())

	v1.POST("/users", func(c *gin.Context) { userController.Create(c) })
	v1.GET("/user/:id", func(c *gin.Context) { userController.Show(c) })
	v1.DELETE("/user/:id", func(c *gin.Context) { userController.Delete(c) })
	v1.PUT("/user/:id", func(c *gin.Context) { userController.Update(c) })

	log.Fatal(engine.Run(":8080"))
}
