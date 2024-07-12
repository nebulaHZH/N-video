package user

import (
	"encoding/json"
	"strconv"

	user "N-video/models"

	"github.com/gin-gonic/gin"
)

var (
	person user.User
)

func Get_user_info_handler(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("uid"))
	u, err := person.GetUserInfo(uid)
	if err == nil {
		c.JSON(200, u)
	} else {
		c.JSON(400, gin.H{"error": "user not found"})
	}
}

func Updae_user_info_handler(c *gin.Context) {
	json.Unmarshal([]byte(c.PostForm("user")), &person)
	state(person.UpdateUserInfo(person), c)
}

func Delete_user_info_handler(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Param("uid"))
	res := person.DeleteUser(uid)
	state(res, c)

}

func Get_video_folder_handler(c *gin.Context) {

}

func state(res bool, c *gin.Context) {
	if res {
		c.JSON(200, gin.H{"status": "success"})
	} else {
		c.JSON(400, gin.H{"status": "fail"})
	}
}
