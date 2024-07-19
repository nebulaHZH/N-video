package user

import (
	"N-video/service/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

// get author information
func Get_author_info(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("uid"))
	c.JSON(200, user.GetAuthorInfo(uid))
}

// controller: get the folders of the user
func Get_video_folder_handler(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("uid"))
	c.JSON(200, user.GetFolderList(uid))
}
