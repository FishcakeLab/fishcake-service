package activity_service

import (
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/database/activity"
)

type ActivityInfoService interface {
	List(pageNum, pageSize int) ([]activity.ActivityInfo, int)
}

type activityInfoService struct {
	Db *database.DB
}

func NewActivityInfoService(db *database.DB) ActivityInfoService {
	return &activityInfoService{Db: db}
}

func (s *activityInfoService) List(pageNum, pageSize int) ([]activity.ActivityInfo, int) {
	infos, count := s.Db.ActivityInfoDB.ListActivityInfo(pageNum, pageSize)
	return infos, count
}
