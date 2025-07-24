package main

import routes "task_manager/delivery/routers"

func main() {
	router := routes.New()
	router.RegisterRoutes()
	router.Run(":8080")
}
