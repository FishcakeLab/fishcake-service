package activity_service

import (
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/database/activity"
)

type ActivityInfoService interface {
	ActivityInfoList(activityFilter, businessAccount, activityStatus, businessName, tokenContractAddr, latitude, longitude, scope string,
		activityId,
		pageNum, pageSize int) ([]activity.ActivityInfo, int)
	ActivityInfo(activityId int) activity.ActivityInfo
	GetActivityRank(monthFilter bool) ([]activity.BusinessRank, error)
}

type activityInfoService struct {
	Db *database.DB
}

func NewActivityInfoService(db *database.DB) ActivityInfoService {
	return &activityInfoService{Db: db}
}

func (s *activityInfoService) ActivityInfoList(activityFilter, businessAccount, activityStatus, businessName, tokenContractAddr, latitude, longitude, scope string, activityId, pageNum, pageSize int) ([]activity.ActivityInfo, int) {
	infos, count := s.Db.ActivityInfoDB.ActivityInfoList(activityFilter, businessAccount, activityStatus, businessName, tokenContractAddr, latitude, longitude, scope, activityId, pageNum, pageSize)
	return infos, count
}

func (s *activityInfoService) ActivityInfo(activityId int) activity.ActivityInfo {
	infos := s.Db.ActivityInfoDB.ActivityInfo(activityId)
	return infos
}

func (s *activityInfoService) GetActivityRank(monthFilter bool) ([]activity.BusinessRank, error) {
	ranks, err := s.Db.ActivityInfoDB.GetActivityRank(monthFilter)
	if err != nil {
		return nil, err
	}
	return ranks, nil
}
