package activity_info

import (
	"context"
	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/database"
	"log"
)

type activityInfoService struct {
	Db *database.DB
}

var ActivityInfoService *activityInfoService

func init() {
	ctx := context.Background()
	cfg, err := config.New(ctx.String("api"))
	if err != nil {
		log.Fatalf("Failed to load config", "err", err)
	}
	log.Printf("Run api start", "HttpHostHost", cfg.HttpHost, "HttpHostHost", cfg.HttpPort)
	db, err := database.NewDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database", "err", err)
	}
	defer func(db *database.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)
	ActivityInfoService = &activityInfoService{Db: db}
}

func (s *activityInfoService) List(pageSize, pageNum string) {

	s.Db.ActivityInfoDB.ListActivityInfo()
}
