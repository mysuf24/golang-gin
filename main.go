package main

import (
	"golang-gin/controllers"
	"golang-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	//inisialisasi gin
	routes := gin.Default()

	//panggil koneksi database
	models.ConnectDataBase()

	//membuat route
	routes.GET("/", func(c *gin.Context) {

		//return response JSON

		c.JSON(200, gin.H{
			"message": "HelloWorld",
		})
	})

	// membuat route get all posts
	routes.GET("/api/posts", controllers.FindPost)

	//membuat route store post
	routes.POST("/api/posts", controllers.StorePost)

	//membuat route detail post
	routes.GET("/api/posts/:id", controllers.FindPostById)

	//membuat route update post
	routes.PUT("/api/posts/:id", controllers.UpdatePost)

	//membuat route delete post
	routes.DELETE("/api/posts/:id", controllers.DeletePost)

	//route server
	routes.Run(":3000")
}
