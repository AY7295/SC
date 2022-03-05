package config

import (
	"SchoolCat/database"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

type Conf struct {
	Util struct {
		AppMode  string `yaml:"AppMode"`
		HttpPort string `yaml:"HttpPort"`
	}
	Mysql struct {
		MySQLHost     string `yaml:"mysqlHost"`
		MySQLPort     string `yaml:"mysqlPort"`
		MySQLUser     string `yaml:"mysqlUser"`
		MySQLPassword string `yaml:"mysqlPWD"`
		DBName        string `yaml:"mysqlDbName"`
	}
	Redis struct {
		RedisHost     string `yaml:"redisHost"`
		RedisPort     string `yaml:"redisPort"`
		RedisPassword string `yaml:"redisPWD"`
		RDBCode       int    `yaml:"redisDbName"`
	}
	Token struct {
		JwtKey       string        `yaml:"jwtKey"`
		ExpireMinute time.Duration `yaml:"expireMinute"`
	}
	Others struct {
		AuthKey string `yaml:"authKey"`
		//ORMLoggerCode logger.LogLevel `yaml:"ORMCode"`
	}
}

var C Conf

func init() {
	var c Conf
	Config(&c)
	path := c.Mysql.MySQLUser + ":" + c.Mysql.MySQLPassword + "@tcp(" + c.Mysql.MySQLHost + ":" + c.Mysql.MySQLPort + ")/" + c.Mysql.DBName + "?charset=utf8mb4&parseTime=true&loc=Local"
	//fmt.Println(path)
	database.MySQL(path)
	database.Redis(c.Redis.RedisHost, c.Redis.RedisPort, c.Redis.RedisPassword, c.Redis.RDBCode)
	//fmt.Println(c.Util)
	C = c

}

func Config(c *Conf) {

	FileByte, err := ioutil.ReadFile("./config/conf.yaml")

	if err != nil {
		log.Printf("配置文件读取错误 err : #%v ", err)
	}

	err = yaml.Unmarshal(FileByte, c)
	if err != nil {
		log.Fatalf("反序列化错误 : %v", err)
	}
}
