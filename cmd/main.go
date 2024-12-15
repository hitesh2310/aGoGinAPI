package main

import (
	"main/pkg/appLogs"
	"main/pkg/router"
)

func init() {
	appLogs.SetUpLogs()
}

func main() {
	s := router.NewServer()
	s.Run("localhost:9090")
}
