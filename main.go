package main

import (
	"gin-fleamarket/models"

	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "item1", Price: 1000, Description: "This is item1", SoldOut: false},
		{ID: 2, Name: "item2", Price: 2000, Description: "This is item2", SoldOut: true},
		{ID: 3, Name: "item3", Price: 3000, Description: "This is item3", SoldOut: false},
	}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("localhost:8080") // 0.0.0.0:8080 でサーバーを立てます。
}
