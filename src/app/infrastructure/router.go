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
	// firebaseAuth := setupFirebase()
	// engine.Use(func(c *gin.Context) {
	// 	c.Set("firebaseAuth", firebaseAuth)
	// })
	v1 := engine.Group("/v1")
	userController := controllers.NewUserController(NewSqlHandler())
	{
		v1.POST("/users", userController.Create)
		// v1.GET("/user/:id", controllers.Show)
		// v1.DELETE("/user/:id", controllers.Delete)
	}
	log.Fatal(engine.Run(":8080"))
}
