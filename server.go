package main

import (
	"github.com/ahmadirfaan/homework4/database"
	"github.com/ahmadirfaan/homework4/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	//Connect database dengan gorm
	database.Connect()

	//Listening Go Fiber
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":8000")
}
