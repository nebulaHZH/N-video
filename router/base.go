package router

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		// 用户信息模块
		UserRoute(v1)
		//视频信息模块
		VideoRoute(v1)
		//上传模块
		UploadRoute(v1)
	}
}
