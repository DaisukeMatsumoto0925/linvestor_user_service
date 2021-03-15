package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	userEngine := engine.Group("/user")
	{
		v1 := userEngine.Group("/v1")
		{
			v1.GET("/", public)
			// v1.GET("/public", private)
		}
	}
	log.Fatal(engine.Run(":8080"))
}

func public(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// func private(c echo.Context) error {

// 	ctx := c.Request().Context()
// 	if ctx == nil {
// 		ctx = context.Background()
// 	}

// 	opt := option.WithCredentialsFile("GOOGLE_APPLICATION_CREDENTIALS")
// 	app, err := fire.NewApp(context.Background(), nil, opt)
// 	if err != nil {
// 		fmt.Printf("error: %v\n", err)
// 		os.Exit(1)
// 	}
// 	auth, err := app.Auth(context.Background())
// 	if err != nil {
// 		fmt.Printf("error: %v\n", err)
// 		os.Exit(1)
// 	}

// 	// クライアントから送られてきた JWT 取得
// 	authHeader := c.Request().Header.Get("Authorization")
// 	idToken := strings.Replace(authHeader, "Bearer ", "", 1)

// 	// JWT の検証
// 	token, err := auth.VerifyIDToken(context.Background(), idToken)
// 	if err != nil {
// 		u := fmt.Sprintf("error verifying ID token: %v\n", err)
// 		return c.JSON(http.StatusBadRequest, u)
// 	}
// 	uid := token.Claims["user_id"]

// 	u := fmt.Sprintf("public is success. uid is %s", uid)

// 	return c.JSON(http.StatusOK, u)
// }
