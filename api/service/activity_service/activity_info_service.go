package activity_service

import (
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/database/activity"
)

type ActivityInfoService interface {
	List(pageSize, pageNum string) ([]activity.ActivityInfo, int64)
}

type activityInfoService struct {
	Db *database.DB
}

func NewActivityInfoService(db *database.DB) ActivityInfoService {
	return &activityInfoService{Db: db}
}

func (s *activityInfoService) List(pageSize, pageNum string) ([]activity.ActivityInfo, int64) {
	infos, count := s.Db.ActivityInfoDB.ListActivityInfo(0, 10)
	return infos, count
}
