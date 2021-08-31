package controllers

import (
	"strings"

	"github.com/ahmadirfaan/homework4/database"
	"github.com/ahmadirfaan/homework4/models"
	"github.com/gofiber/fiber/v2"
)

func GetMovies(c *fiber.Ctx) error {
	slug := c.Params("slug")

	var movies models.Movies
	database.DB.Where("slug = ?", slug).First(&movies)

	if movies.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  "Not Found",
			"result": "Movies Not Found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  nil,
		"result": movies,
	})
}

func CreateMovies(c *fiber.Ctx) error {
	var movies models.Movies
	if err := c.BodyParser(&movies); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"result": "There is a error with your data",
		})
	}
	if movies.Title == "" || movies.Slug == "" || movies.Description == "" || movies.Duration < 0 || movies.Image == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Bad Request",
			"result": "Title, Slug, Description, Image is required And Duration must > 0",
		})
	}
	movies.Slug = strings.TrimSpace(movies.Slug)
	database.DB.Create(&movies)

	if movies.Id == 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":  nil,
			"result": "Movies already exists!",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":  nil,
		"result": movies,
	})
}

func UpdateMovies(c *fiber.Ctx) error {
	slug := c.Params("slug")
	movies := models.Movies{
		Slug: slug,
	}
	if err := c.BodyParser(&movies); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"result": "There is a error with your data",
		})
	}
	movies.Slug = strings.TrimSpace(movies.Slug)
	if movies.Title == "" || movies.Slug == "" || movies.Description == "" || movies.Duration < 0 || movies.Image == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Bad Request",
			"result": "Title, Slug, Description, Image is required And Duration must > 0",
		})
	}

	var findMovies models.Movies
	database.DB.Where("slug = ?", slug).First(&findMovies)
	findMovies.Slug = strings.TrimSpace(findMovies.Slug)
	if movies.Slug == findMovies.Slug {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":  nil,
			"result": "Movies already exists!",
		})
	} else if findMovies.Id == 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":  nil,
			"result": "Movies is not exists!",
		})
	}
	database.DB.Model(&movies).Where("id = ?", findMovies.Id).Updates(&movies)
	movies.Id = findMovies.Id
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":  nil,
		"result": movies,
	})
}

func DeleteMovies(c *fiber.Ctx) error {
	slug := c.Params("slug")
	var movies models.Movies

	database.DB.Where("slug = ?", slug).First(&movies)

	if movies.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  "Not Found",
			"result": "Movies Not Found",
		})
	}

	database.DB.Delete(&movies)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  nil,
		"result": "success",
	})
}
