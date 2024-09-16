package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Article struct {
	gorm.Model
	Title string
	Slug  string `gorm:"unique_index"`
	Desc  string `sql:"type:text;"`
}

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/learngin?charset=utf8&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed conection to database")
	}

	db = database

	database.AutoMigrate(&Article{})

	router := gin.Default()

	v1 := router.Group("/api/v1/")
	{
		articles := v1.Group("/article")
		{
			articles.GET("/", getHome)
			articles.GET("/:slug", getArticle)
			articles.POST("/", postArticle)
		}
	}

	router.Run(":8080")
}

func getHome(c *gin.Context) {
	items := []Article{}
	db.Find(&items)

	c.JSON(200, gin.H{
		"Satus": "Berhasil",
		"Data":  items,
	})
}

func getArticle(c *gin.Context) {
	slug := c.Param("slug")

	var item Article

	result := db.First(&item, "slug = ?", slug)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"Status": "Article not found"})
		} else {
			c.JSON(500, gin.H{"error": result.Error.Error()})
		}
		return
	}

	c.JSON(200, gin.H{
		"Satus":   "Berhasil",
		"Message": item,
	})
}

func postArticle(c *gin.Context) {
	item := Article{
		Title: c.PostForm("title"),
		Desc:  c.PostForm("desc"),
		Slug:  slug.Make(c.PostForm("title")),
	}

	db.Create(&item)

	c.JSON(200, gin.H{
		"Status": "Berhasil Post",
		"Data":   item,
	})
}
