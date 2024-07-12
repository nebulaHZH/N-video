package models

type Favorite struct {
	Uid            int    `json:"uid"`            //用户id
	FavoriteVideos string `json:"favoriteVideos"` //喜欢的视频
}
