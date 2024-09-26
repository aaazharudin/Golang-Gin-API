package routes

import (
	"Learn-Gin/config"
	"Learn-Gin/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

func GetHome(c *gin.Context) {
	items := []models.Article{}
	config.DB.Find(&items)

	c.JSON(200, gin.H{
		"Satus": "Berhasil",
		"Data":  items,
	})
}

func GetProfile(c *gin.Context) {
	var user models.User
	user_id := int(c.MustGet("jwt_user_id").(float64))

	item := config.DB.Where("id = ?", user_id).Preload("Articles", "user_id = ?", user_id).Find(&user)

	c.JSON(200, gin.H{
		"status": "Berhasil ke halaman profile",
		"Data":   item,
	})
}

func GetArticle(c *gin.Context) {
	slug := c.Param("slug")

	var item models.Article

	result := config.DB.First(&item, "slug = ?", slug)
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

func PostArticle(c *gin.Context) {

	slug := slug.Make(c.PostForm("title"))

	// Generate unique slug (efficient loop with counter)
	for {
		var item models.Article
		result := config.DB.First(&item, "slug = ?", slug)

		if result.Error == nil {
			// Slug already exists, append a counter with zero padding
			slug = slug + " - " + strconv.FormatInt(time.Now().Unix(), 10)
		} else if result.Error == gorm.ErrRecordNotFound {
			break
		} else {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}
	}

	item := models.Article{
		Model:  gorm.Model{},
		Title:  c.PostForm("title"),
		Slug:   slug,
		Desc:   c.PostForm("desc"),
		Tag:    c.PostForm("tag"),
		UserID: uint(c.MustGet("Jwt_user_id").(float64)),
	}

	config.DB.Create(&item)

	c.JSON(200, gin.H{
		"Status": "Berhasil Post",
		"Data":   item,
	})
}

func GetArticleByTag(c *gin.Context) {
	tag := c.Param("tag")
	item := []models.Article{}

	config.DB.Where("tag LIKE ?", "%"+tag+"%").Find(&item)

	c.JSON(200, gin.H{
		"status": "Berhasil",
		"Data":   item,
	})
}

func UpdateArticle(c *gin.Context) {
	id := c.Param("id")

	var item models.Article

	result := config.DB.First(&item, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"Status": "Article not found"})
		} else {
			c.JSON(500, gin.H{"error": result.Error.Error()})
		}
		return
	}

	// Update hanya jika id sama
	config.DB.Model(&item).Where("id = ?", id).Updates(models.Article{
		Title: c.PostForm("title"),
		Desc:  c.PostForm("desc"),
		Tag:   c.PostForm("tag"),
	})

	//pembatas agar user hanya bisa mengupdate artikelnya sendiri

	/* 	if uint(c.MustGet("jwt_user_id").(float64)) != item.UserID {
		c.JSON(403, gin.H{
			"status":  "Error",
			"Message": "this data is forbiden",
		})
		c.Abort()
		return
	} */

	c.JSON(200, gin.H{
		"Satus":   "Berhasil",
		"Message": item,
	})
}
