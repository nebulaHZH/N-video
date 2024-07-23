package router

import (
	"N-video/api/user"

	"github.com/gin-gonic/gin"
)

func UserRoute(root *gin.RouterGroup) {
	u := root.Group("/user")
	{
		u.GET("/get", user.Get_user_info_handler)
		u.POST("/modify", user.Updae_user_info_handler)
		u.DELETE("/delete", user.Delete_user_info_handler)
		u.GET("/videofolder", user.Get_video_folder_handler)
		u.POST("/createFolder", user.Create_folder_handler)
		u.POST("/login", user.Login_handler)
		u.POST("/register", user.Rregister_handler)
		u.POST("/modifyPassword", user.Modify_password_handler)
		u.GET("/followed", user.Find_followed)
	}
}
