package database

import (
	"SchoolCat/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
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
