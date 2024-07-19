package history

import "N-video/models"

func GetBrowserHistory(uid int) []models.VideoImpl {
	history = models.HistoryImpl{Uid: uid}
	res := history.Get()
	var v_list []models.VideoImpl
	for _, v := range res {
		video := models.VideoImpl{Vid: v.Vid}
		v_list = append(v_list, video.GetVideo())
	}
	return v_list
}

func DeleteBrowserHistory(uid int, vid_list []string) bool {
	for _, vid := range vid_list {
		history = models.HistoryImpl{Uid: uid, Vid: vid}
		history.Delete()
	}
	return len(GetBrowserHistory(uid)) == 0
}
