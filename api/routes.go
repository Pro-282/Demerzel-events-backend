package api

import (
	"demerzel-events/internal/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

func BuildRoutesHandler() *gin.Engine {
	r := gin.New()

	if os.Getenv("APP_ENV") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	r.GET("/health", handlers.HealthHandler)

	r.GET("/api/users/profile/:id", handlers.HandleGetUserID)
	r.PUT("/api/users/profile/:id", handlers.HandleUpdateUser)

	


	// OAuth routes
	authRoute := r.Group("/oauth")
	authRoute.GET("/initialize", handlers.InitalizeOAuthSignIn)
	authRoute.GET("/callback", handlers.HandleOAuthCallBack)

	return r
}
