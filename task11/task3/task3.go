package main

import (
	"github.com/gin-gonic/gin"
)

type Request struct {
	A int `json:"a"`
	B int `json:"b"`
}


func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		name := c.Query("name")
		c.JSON(200, gin.H{
			"message": "Привет, " + name + "!",
		})
	})

	r.POST("/divide", func(c *gin.Context) {
		var req Request
		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Неверный формат JSON"})
			return
		}
		if req.B == 0{
			c.JSON(400, gin.H{"error": "деление на ноль"})
			return
		}
		c.JSON(200, gin.H{
			"result": req.A/req.B,
		})
		
	})

	r.Run(":8080")
}
