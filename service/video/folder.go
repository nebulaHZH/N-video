package video

import "N-video/models"

func AddVideoToFolder(fid string, vid string) bool {
	relation = models.Relation{Fid: fid, Vid: vid}
	relation.AddIntoVideoFolder(fid, vid)
	return true
}

func RemoveVideoFromFolder(fid string, vid string) bool {
	relation = models.Relation{Fid: fid, Vid: vid}
	relation.RemoveFromVideoFolder(fid, vid)
	return true
}
