package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zvihell/his/config"
	"github.com/zvihell/his/controllers"
)

func main() {
	config.InitialiseDBConnection()

	r := gin.Default()
	r.POST("api/setInfo", controllers.CreatePost)
	r.GET("api/getInfo", controllers.GetPosts)

	r.Run(":8888")
}
