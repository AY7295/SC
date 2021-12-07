package main

import (
	"SchoolCat/database"
	"SchoolCat/router"
	"fmt"
)


func main() {
	database.Link()
	engine := router.Router()
	if err := engine.Run(":7000"); err != nil {
		fmt.Println(err)
	}
}
