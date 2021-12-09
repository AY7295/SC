
package database

import (
	"SchoolCat/model"
	"gorm.io/driver/mysql"
"gorm.io/gorm"
"gorm.io/gorm/schema"
"log"
)

const (
	userName = "root"
	password = "12345qwert"
	ip = "127.0.0.1"
	port = "3306"
	dbName = "schoolcat"
)


func Link() *gorm.DB {
	var (
		DB *gorm.DB
		err error
	)
	path := userName+ ":"+ password+"@tcp("+ip+ ":"+ port+")/"+ dbName+ "?charset=utf8&parseTime=true&loc=Local"
	DB,err =gorm.Open(mysql.New(mysql.Config{
		DSN:path ,
		DefaultStringSize: 256,
	}), &gorm.Config{
		NamingStrategy:schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Panic(err.Error())
	}
	//根据model创建一个表
	err = DB.AutoMigrate(&model.User{},&model.Admin{},&model.Share{},&model.ShareImage{},&model.UserComment{})

	if err != nil {
		log.Panic(err.Error())
	}
	return DB
}

