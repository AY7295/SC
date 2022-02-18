package database

import (
	"SchoolCat/config"
	"SchoolCat/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func MySQL(path string) {

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               path,
		DefaultStringSize: 256,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: MysqlLogger,
	})
	if err != nil {
		log.Println("mysql数据库链接失败：" + err.Error())
		return
	} else {
		fmt.Println("mysql数据库链接成功")
	}

	//根据model创建一个表
	err = db.AutoMigrate(&model.User{}, &model.Admin{}, &model.Share{}, &model.ShareImage{}, &model.UserComment{}, &model.CatCard{}, &model.CatCardSrc{}, &model.CatCardComment{}, &model.Tip{}, &model.TipComment{}, &model.TipCommentLike{}, &model.TipSrc{}, &model.ShareLike{}, &model.ShareCommentLike{})
	if err != nil {
		log.Panic("建表失败：" + err.Error())
		return
	} else {
		fmt.Println("建表成功")
	}

	DB = db
}

var MysqlLogger logger.Interface = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
	logger.Config{
		SlowThreshold:             time.Nanosecond,                       // 快 SQL 阈值
		LogLevel:                  logger.LogLevel(config.ORMLoggerCode), // 日志级别
		IgnoreRecordNotFoundError: false,                                 // ErrRecordNotFound（记录未找到）错误
		Colorful:                  true,                                  // 用彩色打印
	},
)
