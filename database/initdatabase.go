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
var (
	DB  *gorm.DB
	err error
)
func Link() *gorm.DB {

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
	}else {fmt.Println("数据库链接成功")}

	//根据model创建一个表
	err = DB.AutoMigrate(&model.User{}, &model.Admin{}, &model.Share{}, &model.ShareImage{}, &model.UserComment{}, &model.CatCard{}, &model.CatCardSrc{}, &model.CatCardComment{}, &model.Tip{},  &model.TipComment{},&model.TipCommentLike{},&model.TipSrc{},&model.ShareLike{},&model.ShareCommentLike{})
	if err != nil {
		log.Panic(err.Error())
	}else {fmt.Println("建表成功")}
	return DB
}
