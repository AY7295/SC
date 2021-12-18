package main

import (
	"SchoolCat/database"
	"SchoolCat/router"
	"fmt"
)

func main() {
	fmt.Println("SchoolCat")
	database.Link()
	router.Router()
}
