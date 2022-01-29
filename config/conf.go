package config

import (
	"SchoolCat/database"
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	JwtKey  string
	AuthKey string
)

func init() {

	file, err := ini.Load("./config/conf.ini")
	if err != nil {
		fmt.Println("配置文件读取错误", err)
	}
	LoadSever(file)
	LoadData(file)
	LoadKey(file)
	path := DBUser + ":" + DBPassword + "@tcp(" + DBHost + ":" + DBPort + ")/" + DBName + "?charset=utf8&parseTime=true&loc=Local"
	database.Link(path)
}

func LoadSever(file *ini.File) {
	AppMode = file.Section("sever").Key("AppMode").MustString("debug")
	HttpPort = file.Section("sever").Key("HttpPort").MustString(":8080")

}

func LoadData(file *ini.File) {

	DBHost = file.Section("database").Key("dbHost").MustString("localhost")
	DBPort = file.Section("database").Key("dbPort").MustString("3306")
	DBUser = file.Section("database").Key("dbUser").MustString("root")
	DBPassword = file.Section("database").Key("dbPWD").MustString("nil")
	DBName = file.Section("database").Key("dbName").MustString("test")

}
func LoadKey(file *ini.File) {
	JwtKey = file.Section("KEY").Key("jwtKey").MustString("helloWorld")
	AuthKey = file.Section("KEY").Key("authKey").MustString("12345")
}
