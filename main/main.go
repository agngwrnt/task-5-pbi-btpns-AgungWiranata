package main

import (
	"btpn/database"
	"btpn/router"
)

func main() {
	database.ConnectDB()
	r := router.SetupRouter()
	r.Run(":8080")
}
