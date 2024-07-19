package router

import (
	"N-video/api/upload"

	"github.com/gin-gonic/gin"
)

func UploadRoute(root *gin.RouterGroup) {
	up := root.Group("/upload")
	{
		up.POST("/video", upload.Upload_video_handler)
		up.POST("/cover", upload.Upload_cover_handler)
		up.POST("/avatar", upload.Upload_avatar_handler)
	}
}
