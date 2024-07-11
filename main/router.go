package main

import (
	user "N-video/models"
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	person user.User
)

func router(router *gin.Engine) *gin.Engine {
	v1 := router.Group("/api/v1")
	{
		user_info := v1.Group("/user_info")
		{
			user_info.GET("/userinfo", get_user_info_handler)
			user_info.POST("/modifyInfo", updae_user_info_handler)
			user_info.DELETE("/:id", delete_user_info_handler)
		}
	}
	return router
}
func get_user_info_handler(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("uid"))
	u, err := person.GetUserInfo(uid)
	if err == nil {
		c.JSON(200, u)
	} else {
		c.JSON(400, gin.H{"error": "user not found"})
	}
}

func updae_user_info_handler(c *gin.Context) {
	json.Unmarshal([]byte(c.PostForm("user")), &person)
	state(person.UpdateUserInfo(person), c)
}

func delete_user_info_handler(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Param("uid"))
	res := person.DeleteUser(uid)
	state(res, c)

}

func state(res bool, c *gin.Context) {
	if res {
		c.JSON(200, gin.H{"status": "success"})
	} else {
		c.JSON(400, gin.H{"status": "fail"})
	}
}
