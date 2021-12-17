package database

import (
	"SchoolCat/config"
	"SchoolCat/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

func Link() *gorm.DB {
	var (
		DB  *gorm.DB
		err error
	)
	path := config.DBUser + ":" + config.DBPWD + "@tcp(" + config.DBHost + ":" + config.DBPort + ")/" + config.DBName + "?charset=utf8&parseTime=true&loc=Local"
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               path,
		DefaultStringSize: 256,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	//根据model创建一个表
	err = DB.AutoMigrate(&model.User{}, &model.Admin{}, &model.Share{}, &model.ShareImage{}, &model.UserComment{}, &model.CatCard{}, &model.CatCardSrc{}, &model.CatCardComment{}, &model.Tip{},  &model.TipComment{},&model.TipCommentLike{},&model.TipSrc{},&model.ShareLike{},&model.ShareCommentLike{})

	if err != nil {
		log.Panic(err.Error())
	}
	return DB
}
