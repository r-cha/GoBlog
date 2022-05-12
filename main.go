package main

import (
	"r-cha/goblog/controllers"
	"r-cha/goblog/db"
)

func main() {
	db.Connect()

	r := controllers.NewRoutes()

	r.Run()
}
