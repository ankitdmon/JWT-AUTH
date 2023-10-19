package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.con/ankitdmon/initializers"
)

func init() {
	initializers.LoadENV()
	initializers.ConnectDB()
}

func main() {
	fmt.Println("Hello Golang!!")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()	
}
