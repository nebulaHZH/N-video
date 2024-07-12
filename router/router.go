package router

import (
	comment "N-video/service/home"
	video "N-video/service/home"
	upload "N-video/service/upload"
	log "N-video/service/user"
	userinfo "N-video/service/user"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) *gin.Engine {
	v1 := router.Group("/api/v1")
	{
		// 用户信息模块
		user_info := v1.Group("/user_info")
		{
			user_info.GET("/get", userinfo.Get_user_info_handler)
			user_info.POST("/modify", userinfo.Updae_user_info_handler)
			user_info.DELETE("/delete", userinfo.Delete_user_info_handler)
			user_info.GET("/videofolder", userinfo.Get_video_folder_handler)
		}
		//登录注册模块
		log_register := v1.Group("/log")
		{
			// log_register.POST("/login", log.Login_handler)
			log_register.POST("/register", log.Rregister_handler)
			log_register.POST("/modifyPassword", log.Modify_password_handler)
		}
		//视频信息模块
		video_info := v1.Group("/video")
		{
			video_info.POST("/get", video.Get_video_info_handler)
		}
		//评论模块
		review := v1.Group("/comment")
		{
			review.GET("/add", comment.Add_comment_handler)
			review.POST("/get", comment.Get_comment_handler)
			review.GET("/delete", comment.Delete_comment_handler)
		}
		//上传模块
		upload_file := v1.Group("/upload")
		{
			upload_file.POST("/video", upload.Upload_video_handler)
			upload_file.POST("/cover", upload.Upload_cover_handler)
			upload_file.POST("/avatar", upload.Upload_avatar_handler)
		}
	}
	return router
}
