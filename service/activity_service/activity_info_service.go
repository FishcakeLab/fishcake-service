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
	GetActivityRank(monthFilter bool) ([]MiningRankResponse, error)
}

type activityInfoService struct {
	Db *database.DB
}

// MiningRankResponse — 商家挖矿排行榜返回结构
type MiningRankResponse struct {
	Rank            int    `json:"rank"`            // 排名序号（1开始）
	BusinessAccount string `json:"businessAccount"` // 商家账户地址
	TotalMined      string `json:"totalMined"`      // 累计挖矿产出（单位：Wei）
	Month           bool   `json:"month"`           // 是否为月榜 (true=本月, false=总榜)
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

func (s *activityInfoService) GetActivityRank(monthFilter bool) ([]MiningRankResponse, error) {
	ranks, err := s.Db.ActivityInfoDB.GetActivityRank(monthFilter)
	if err != nil {
		return nil, err
	}
	var res []MiningRankResponse
	for i, item := range ranks {
		res = append(res, MiningRankResponse{
			Rank:            i + 1,
			BusinessAccount: item.BusinessAccount,
			TotalMined:      item.TotalMined,
			Month:           monthFilter,
		})
	}
	return res, nil
}
