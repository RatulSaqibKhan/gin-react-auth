package app

import (
	"gin-react-auth/app/controllers/users"
	"time"

	"github.com/gin-contrib/cors"
)

func MapUrls() {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8665"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.POST("api/register", users.Register)
	router.POST("api/login", users.Login)
	router.GET("api/home", users.Home)
	router.GET("api/logout", users.Logout)
}
