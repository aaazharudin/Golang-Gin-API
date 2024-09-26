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
		v1.GET("/profile", middleware.IsAuth(), routes.GetProfile)
		v1.GET("/check", middleware.IsAuth(), routes.CheckToken)

		v1.GET("/article/:slug", routes.GetArticle)
		articles := v1.Group("/articles")
		{
			articles.GET("/", routes.GetHome)
			articles.POST("/", middleware.IsAuth(), routes.PostArticle)
			articles.GET("/tag/:tag", routes.GetArticleByTag)
			articles.PUT("/update/:id", middleware.IsAuth(), routes.UpdateArticle)
		}
	}

	router.Run(":8080")

}
