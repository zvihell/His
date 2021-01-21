package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zvihell/his/config"
)

type Post struct {
	ID   int64  `json:"id"`
	Info string `json:"info"`
}

func CreatePost(c *gin.Context) {
	var reqBody Post
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": "Invalid request body",
		})
		return
	}
	res, _ := config.DBClient.Exec("INSERT INTO his(info) VALUES(?);",
		reqBody.Info,
	)

	id, _ := res.LastInsertId()

	c.JSON(http.StatusCreated, gin.H{
		"status": "OK",
		"id":     id,
	})
}

func GetPosts(c *gin.Context) {
	var posts []Post

	rows, err := config.DBClient.Query("SELECT id,info FROM his;")

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": true,
		})
		return
	}

	for rows.Next() {
		var singlePost Post
		if err := rows.Scan(&singlePost.ID, &singlePost.Info); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": true,
			})
			return
		}
		posts = append(posts, singlePost)
	}
	c.JSON(http.StatusOK, posts)
}
