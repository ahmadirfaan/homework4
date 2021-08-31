package routes

import (
	"github.com/ahmadirfaan/homework4/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("movie/:slug", controllers.GetMovies)
	app.Post("movie", controllers.CreateMovies)
	app.Put("movie/:slug", controllers.UpdateMovies)
	app.Delete("movie/:slug", controllers.DeleteMovies)
}