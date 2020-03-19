package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("oauth2/v4/token", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "1",
		})
	})

	r.GET("/tokeninfo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "1",
		})
	})

	_ = r.Run(":3001")
}
