package activity_service

import (
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/database/activity"
)

type ActivityInfoExtService interface {
	ActivityInfoExtList(pageNum, pageSize int) ([]activity.ActivityInfoExt, int)
	ActivityInfoExt(activityId int) activity.ActivityInfoExt
}

type activityInfoExtService struct {
	Db *database.DB
}

func NewActivityInfoExtService(db *database.DB) ActivityInfoExtService {
	return &activityInfoExtService{Db: db}
}

func (s *activityInfoExtService) ActivityInfoExtList(pageNum, pageSize int) ([]activity.ActivityInfoExt, int) {
	infos, count := s.Db.ActivityInfoExtDB.ActivityInfoExtList(int(uint64(pageNum)), pageSize)
	return infos, count
}

func (s *activityInfoExtService) ActivityInfoExt(activityId int) activity.ActivityInfoExt {
	infos := s.Db.ActivityInfoExtDB.ActivityInfoExt(activityId)
	return infos
}
