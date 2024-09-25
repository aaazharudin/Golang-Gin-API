package main

import (
	"Learn-Gin/config"
	"Learn-Gin/middleware"
	"Learn-Gin/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func main() {
	config.InitDB()
	//menutup koneksi database ketika aplikasi telah selesai
	defer func() {
		if sqlDB, err := config.DB.DB(); err == nil {
			sqlDB.Close()
		} else {
			log.Printf("failed to get database instance: %v", err)
		}
	}()

	gotenv.Load()

	//set up routing/router
	router := gin.Default()

	v1 := router.Group("/api/v1/")
	{
		v1.GET("/auth/:provider", routes.RedirectHandler)
		v1.GET("/auth/:provider/callback", routes.CallbackHandler)

		//testing token user
		v1.GET("/check", middleware.IsAuth(), routes.CheckToken)

		articles := v1.Group("/article")
		{
			articles.GET("/", routes.GetHome)
			articles.GET("/:slug", routes.GetArticle)
			articles.POST("/", middleware.IsAuth(), routes.PostArticle)
		}
	}

	router.Run(":8080")

}
