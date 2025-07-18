package main

import "task_manager/router"

func main() {
	router := router.New()
	router.RegisterRoutes()
	router.Run(":8080")
}