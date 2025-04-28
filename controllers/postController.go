package controllers

import (
	"errors"
	"golang-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// validasi post input
type ValidatePostInput struct {
	Title   string  `json: "title" binding: "required"`
	Content string `json: "content" binding: "required"`
}

// error message
type ErrorMsg struct {
	Feild   string `json: "field"`
	Message string `json: "message"`
}

// func get error msg
func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "this field is required"
	}
	return "Unknown error"
}

// get all post
func FindPost(c *gin.Context) {

	// mengambil data dari database menggunakan model
	var posts []models.Post
	models.DB.Find(&posts)

	//return json
	c.JSON(200, gin.H{
		"success": true,
		"message": "List Data Posts",
		"data":    posts,
	})
}

// store a post
func StorePost(c *gin.Context) {
	//validasi input
	var input ValidatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": out})
		}
		return
	}
	//create post
	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
	}

	models.DB.Create(&post)

	//return resonse json
	c.JSON(201, gin.H{
		"success": true,
		"message": "Post Created Successfully",
		"data":    post,
	})
}
// get post by id
func FindPostById(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Detail Data Post By ID : " + c.Param("id"),
		"data":    post,
	})
}

// update post
func UpdatePost(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	//validate input
	var input ValidatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	//update post
	models.DB.Model(&post).Updates(input)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Post Updated Successfully",
		"data":    post,
	})
}
// delete post
func DeletePost(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	//delete post
	models.DB.Delete(&post)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Post Deleted Successfully",
	})
}