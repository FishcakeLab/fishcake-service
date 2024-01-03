package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/utils/mylogs"
)

// Orm 全局变量 ,初始化后的连接
var Orm *orm

type orm struct {
	myDb *gorm.DB
}

func NewOrm(cfg *config.Config) *orm {
	o := new(orm)
	o.initDBConfig(cfg)
	return o
}

func (db *orm) initDBConfig(cfg *config.Config) (err error) {

	if db.myDb != nil {
		return
	}
	password := cfg.DbPassword
	username := cfg.DbUser
	host := cfg.DbHost
	port := cfg.DbPort
	database := cfg.DbName
	dsn := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai"
	dsn = fmt.Sprintf(dsn, host, username, password, database, port)

	// 初始化GORM日志配置
	writer := mylogs.MyLogWriter()
	newLogger := logger.New(
		log.New(writer, "\r\n", log.Ldate|log.Ltime|log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level(这里记得根据需求改一下)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	gormDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	if err != nil {
		log.Fatalf("数据库连接错误：%v", err)
		return err
	}

	// 尝试与数据库建立连接
	sqlDB, err := gormDb.DB()
	if err != nil {
		log.Println(err)
		return err
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Println(err)
		return err
	}
	//设置连接池
	//空闲
	sqlDB.SetMaxIdleConns(10)
	//打开
	sqlDB.SetMaxOpenConns(20)
	db.myDb = gormDb

	log.Println("数据库初始化成功")
	return nil
}

func (db *orm) DB() *gorm.DB {
	return db.myDb
}
