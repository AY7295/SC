package main

import (
	"SchoolCat/config"
	"SchoolCat/router"
	"fmt"
)

func main() {
	fmt.Println("SchoolCat")
	config.Init()
	router.Router()
}
