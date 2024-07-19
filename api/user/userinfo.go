package user

import (
	"strconv"

	"N-video/models"
	"N-video/service/user"

	"github.com/gin-gonic/gin"
)

// 获取用户详细信息处理函数-GET
func Get_user_info_handler(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("uid"))
	/*middleware here*/
	c.JSON(200, user.GetUserInfo(uid))
}

// 更新用户详细信息-POST，JSON对象格式传输
func Updae_user_info_handler(c *gin.Context) {
	p := models.User{}
	c.BindJSON(&p)
	c.JSON(200, user.UpdateUserInfo(p))
}

// 删除用户信息-DELETE
func Delete_user_info_handler(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Param("uid"))
	/*middleware here*/
	u := user.DeleteUserInfo(uid)
	if u.Uid == 0 {
		c.JSON(200, gin.H{"msg": "success"})
	} else {
		c.JSON(400, gin.H{"msg": "fail"})
	}
}

// controller: get the followed users of the user
func Find_followed(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("uid"))
	c.JSON(200, user.GetFollowed(uid))
}

// controller: cancel follow the user
func Cancel_follow(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("uid"))
	fid, _ := strconv.Atoi(c.Query("fid"))
	//convert to service and handle the cancel function
	c.JSON(200, user.CancelFollow(uid, fid))
}

// controller: find the followers of the user
func Find_follower(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("uid"))
	c.JSON(200, user.GetFollower(uid))
}
