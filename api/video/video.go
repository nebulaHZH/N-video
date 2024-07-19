package video

import (
	"N-video/service/video"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Get_video_info_handler(c *gin.Context) {
	// TODO: 获取视频信息
	c.JSON(200, gin.H{
		"message": "get video info",
	})
}

func Get_video_info(c *gin.Context) {
	vid := c.Query("vid")
	c.JSON(200, video.GetVideoInfo(vid))
}

func Upload_video(c *gin.Context) {
	// TODO: 上传视频 + 上传视频集 , 考虑到上传功能下次写
	c.JSON(200, gin.H{
		"message": "upload video",
	})
}

func Delete_video(c *gin.Context) {
	vid := c.Query("vid")
	// uid, _ := strconv.Atoi(c.Query("uid"))
	// 这里要加token 中间件
	c.JSON(200, video.DeleteVideo(vid))
}

// controller: histroy of the user
func Browser_history(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("uid"))
	c.JSON(200, video.GetBrowserHistory(uid))
}

// controller: delete the browser history of the user
func Delete_browser_history(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("uid"))
	vid_list := c.PostFormArray("vid_list[]")
	c.JSON(200, video.DeleteBrowserHistory(uid, vid_list))
}

// controller: like the video
func Like_Video(c *gin.Context) {
	vid := c.Query("vid")
	uid, _ := strconv.Atoi(c.Query("uid"))
	c.JSON(200, video.LikeVideo(uid, vid))
}

// controller: dislike the video
func Dislike_Video(c *gin.Context) {
	vid := c.Query("vid")
	uid, _ := strconv.Atoi(c.Query("uid"))
	c.JSON(200, video.DislikeVideo(uid, vid))
}
