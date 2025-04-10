package main

import (
	"online-store/config"
	"online-store/routes"
)

func main() {
    config.ConnectDB()
    r := routes.SetupRoutes()
    r.Run(":8080")
}
