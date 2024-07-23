package history

import "N-video/models"

func GetBrowserHistory(uid int) []models.Video {
	history = models.History{Uid: uid}
	res := history.Get()
	var v_list []models.Video
	for _, v := range res {
		video := models.Video{Vid: v.Vid}
		v_list = append(v_list, video.GetVideo())
	}
	return v_list
}

func DeleteBrowserHistory(uid int, vid_list []string) bool {
	for _, vid := range vid_list {
		history = models.History{Uid: uid, Vid: vid}
		history.Delete()
	}
	return len(GetBrowserHistory(uid)) == 0
}
