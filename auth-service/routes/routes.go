package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xavierspacelix/logistic-system/auth-service/handler"
	"github.com/xavierspacelix/logistic-system/auth-service/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/auth")
	{
		api.POST("/register", handler.Register)
		api.POST("/login", handler.Login)

		apiAuth := api.Group("/")
		apiAuth.Use(middleware.JWTAuth())
		{
			apiAuth.GET("/me", handler.Me)
		}
	}
}
