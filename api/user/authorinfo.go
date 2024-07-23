package user

import (
	"N-video/service/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

// controller: get the folders of the user
func Get_video_folder_handler(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("uid"))
	c.JSON(200, user.GetFolderList(uid))
}

// controller : create a new folder
func Create_folder_handler(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("uid"))
	folder_name := c.PostForm("folder_name")
	c.JSON(200, user.CreateFolder(uid, folder_name))
}
