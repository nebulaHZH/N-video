package video

import (
	"N-video/models"
	"N-video/service/video"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Add_comment_handler(c *gin.Context) {
	var comment models.Comment
	c.BindJSON(&comment)
	if comment.Vid == "" {
		c.JSON(200, gin.H{
			"message": "Vid is empty",
		})
		return
	}
	res := video.AddComment(comment)
	c.JSON(200, res)
}

func Get_comment_handler(c *gin.Context) {
	vid := c.PostForm("vid")
	c.JSON(200, video.GetCommentList(vid))
}

func Delete_comment_sender(c *gin.Context) {
	cid := c.PostForm("cid")
	uid, _ := strconv.Atoi(c.PostForm("uid"))
	c.JSON(200, video.DeleteCommentBySender(uid, cid))
}
