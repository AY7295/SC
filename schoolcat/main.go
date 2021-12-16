package main

import (
	"SchoolCat/router"
	"fmt"
)


func main() {
	engine := router.Router()
	if err := engine.Run("localhost:7000"); err != nil {
		fmt.Println(err)
	}
}
