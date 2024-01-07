package activity_info

import (
	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/database/activity"
	"log"
)

type activityInfoService struct {
	Db *database.DB
}

var ActivityInfoService *activityInfoService

func init() {
	cfg, err := config.New("./config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config", "err", err)
	}
	log.Printf("Run api start", "HttpHostHost", cfg.HttpHost, "HttpHostHost", cfg.HttpPort)
	db, err := database.NewDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database", "err", err)
	}
	ActivityInfoService = &activityInfoService{Db: db}
}

func (s *activityInfoService) List(pageSize, pageNum string) ([]activity.ActivityInfo, int64) {
	infos, count := s.Db.ActivityInfoDB.ListActivityInfo(0, 10)
	return infos, count
}
