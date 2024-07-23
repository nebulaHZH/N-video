package video

import (
	"N-video/service/history"
	"N-video/service/like"
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
	uid, _ := strconv.Atoi(c.Query("uid"))
	c.JSON(200, video.GetVideoInfo(vid, uid))
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
	res := video.DeleteVideo(vid)
	if res {
		// if you delete the video successfully,then del the comment
		res_c := video.DeleteCommentByVid(vid)
		c.JSON(200, gin.H{
			"del_video":   res,
			"del_comment": res_c,
		})
	} else {
		c.JSON(200, res)
	}
}

// controller: histroy of the user
func Browser_history(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("uid"))
	c.JSON(200, history.GetBrowserHistory(uid))
}

// controller: delete the browser history of the user
func Delete_browser_history(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("uid"))
	vid_list := c.PostFormArray("vid_list[]")
	c.JSON(200, history.DeleteBrowserHistory(uid, vid_list))
}

// controller: like the video
func Like_Video(c *gin.Context) {
	vid := c.Query("vid")
	uid, _ := strconv.Atoi(c.Query("uid"))
	c.JSON(200, like.LikeVideo(uid, vid))
}

// controller: dislike the video
func Dislike_Video(c *gin.Context) {
	vid := c.Query("vid")
	uid, _ := strconv.Atoi(c.Query("uid"))
	c.JSON(200, like.DislikeVideo(uid, vid))
}

func Add_to_folder(c *gin.Context) {
	fid := c.PostForm("fid")
	vid := c.PostForm("vid")
	res := video.AddVideoToFolder(fid, vid)
	c.JSON(200, res)
}

func Remove_from_folder(c *gin.Context) {
	fid := c.PostForm("fid")
	vid := c.PostForm("vid")
	res := video.RemoveVideoFromFolder(fid, vid)
	c.JSON(200, res)
}
