package config

import (
	"SchoolCat/database"
	"fmt"
	"gopkg.in/ini.v1"
	"runtime"
	"strconv"
)

var (
	AppMode  string
	HttpPort string

	MySQLHost     string
	MySQLPort     string
	MySQLUser     string
	MySQLPassword string
	DBName        string

	RedisHost     string
	RedisPort     string
	RedisPassword string
	RDBCode       int

	JwtKey  string
	AuthKey string

	ORMLoggerCode int
)

func init() {

	var (
		file *ini.File
		err  error
	)

	//fmt.Println(runtime.GOOS)

	if runtime.GOOS == "windows" {
		file, err = ini.Load("./config/local-conf.ini")
	} else {
		file, err = ini.Load("./config/server-conf.ini")
	}

	if err != nil {
		fmt.Println("配置文件读取错误", err)
		return
	}
	LoadFile(file)

	path := MySQLUser + ":" + MySQLPassword + "@tcp(" + MySQLHost + ":" + MySQLPort + ")/" + DBName + "?charset=utf8&parseTime=true&loc=Local"
	database.MySQL(path)
	database.Redis(RedisHost, RedisPort, RedisPassword, RDBCode)
}

func LoadFile(file *ini.File) {

	AppMode = file.Section("sever").Key("AppMode").MustString("debug")
	HttpPort = file.Section("sever").Key("HttpPort").MustString(":8080")

	MySQLHost = file.Section("mysql").Key("mysqlHost").MustString("localhost")
	MySQLPort = file.Section("mysql").Key("mysqlPort").MustString("3306")
	MySQLUser = file.Section("mysql").Key("mysqlUser").MustString("root")
	MySQLPassword = file.Section("mysql").Key("mysqlPWD").MustString("nil")
	DBName = file.Section("mysql").Key("mysqlDbName").MustString("test")

	RedisHost = file.Section("redis").Key("redisHost").MustString("localhost")
	RedisPort = file.Section("redis").Key("redisPort").MustString("6379")
	//RedisUser = file.Section("redis").Key("redisUser").MustString("root")
	RedisPassword = file.Section("redis").Key("redisPWD").MustString("nil")
	str1 := file.Section("redis").Key("redisDbName").MustString("test")
	RDBCode, _ = strconv.Atoi(str1)

	JwtKey = file.Section("KEY").Key("jwtKey").MustString("helloWorld")
	AuthKey = file.Section("KEY").Key("authKey").MustString("12345")

	str2 := file.Section("ORM").Key("code").MustString("4")
	ORMLoggerCode, _ = strconv.Atoi(str2)
}
