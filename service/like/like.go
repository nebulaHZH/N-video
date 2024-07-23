package like

import "N-video/models"

var favorite models.FavoriteInter

func LikeVideo(uid int, vid string) bool {
	favorite = models.Favorite{Uid: uid, Vid: vid}
	favorite.Like()
	return favorite.GetLike(vid).Uid != 0
}

func DislikeVideo(uid int, vid string) bool {
	favorite = models.Favorite{Uid: uid, Vid: vid}
	favorite.Dislike()
	// select again and check if success
	return favorite.GetLike(vid).Uid == 0
}
