package home

import "github.com/gin-gonic/gin"

func Add_comment_handler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Add_comment_handler",
	})
}

func Get_comment_handler(c *gin.Context) {

}

func Delete_comment_handler(c *gin.Context) {

}
