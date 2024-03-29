package activity_service

import (
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/database/activity"
)

type ActivityInfoService interface {
	ActivityInfoList(pageNum, pageSize int) ([]activity.ActivityInfo, int)
	ActivityInfo(activityId int) activity.ActivityInfo
}

type activityInfoService struct {
	Db *database.DB
}

func NewActivityInfoService(db *database.DB) ActivityInfoService {
	return &activityInfoService{Db: db}
}

func (s *activityInfoService) ActivityInfoList(pageNum, pageSize int) ([]activity.ActivityInfo, int) {
	infos, count := s.Db.ActivityInfoDB.ActivityInfoList(int(uint64(pageNum)), pageSize)
	return infos, count
}

func (s *activityInfoService) ActivityInfo(activityId int) activity.ActivityInfo {
	infos := s.Db.ActivityInfoDB.ActivityInfo(activityId)
	return infos
}
