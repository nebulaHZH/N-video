package router

import (
	"N-video/api/video"

	"github.com/gin-gonic/gin"
)

func VideoRoute(root *gin.RouterGroup) {
	v := root.Group("/video")
	{
		v.POST("/get", video.Get_video_info_handler)
		v.POST("/upload", video.Upload_video)
		v.POST("/delete", video.Delete_video)
		v.POST("/addtofolder", video.Add_to_folder)
		v.POST("/removefromfolder", video.Remove_from_folder)
		review := v.Group("/comment")
		{
			review.GET("/add", video.Add_comment_handler)
			review.POST("/get", video.Get_comment_handler)
			review.GET("/delete", video.Delete_comment_sender)
		}
	}

}
