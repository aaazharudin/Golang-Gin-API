package main

import (
	"Learn-Gin/config"
	"Learn-Gin/routes"
	"log"

	"github.com/gin-gonic/gin"
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

	//set up routing/router
	router := gin.Default()

	v1 := router.Group("/api/v1/")
	{
		articles := v1.Group("/article")
		{
			articles.GET("/", routes.GetHome)
			articles.GET("/:slug", routes.GetArticle)
			articles.POST("/", routes.PostArticle)
		}
	}

	router.Run(":8080")

}
