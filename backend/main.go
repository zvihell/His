package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zvihell/his/config"
	"github.com/zvihell/his/controllers"
)

func main() {
	config.InitialiseDBConnection()

	r := gin.Default()
	r.POST("/setInfo", controllers.CreatePost)
	r.GET("/getInfo", controllers.GetPosts)

	r.Run(":8888")
}
