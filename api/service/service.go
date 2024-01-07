package service

import (
	"github.com/FishcakeLab/fishcake-service/api/service/activity_service"
	"github.com/FishcakeLab/fishcake-service/database"
)

type service struct {
	ActivityInfoService activity_service.ActivityInfoService
}

var BaseService *service

func NewBaseService(db *database.DB) {
	BaseService = &service{
		ActivityInfoService: activity_service.NewActivityInfoService(db),
	}
}
