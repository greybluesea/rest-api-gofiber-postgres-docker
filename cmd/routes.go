package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/greybluesea/dockerised-fullstack-webapp-gofiber-gorm-postgres/database"
	"github.com/greybluesea/dockerised-fullstack-webapp-gofiber-gorm-postgres/models"
)

func setupRoutes(app *fiber.App) {
	app.Get("/hello", helloHandler)
	app.Get("/", homeHandler)
	app.Post("/create", createHandler)
}

func helloHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World! Tony here 👋!!")
}
func homeHandler(c *fiber.Ctx) error {

	facts := []models.Fact{}
	database.DB.Find(&facts)

	// 	return c.Status(200).JSON(facts)
	return c.Render("index", fiber.Map{"Title": "Fun Facts", "Subtitle": "Dockerised Fullstack WebApp(GoFiber + GORM + Postgres) - learned from Div Rhino"})
}

func createHandler(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(&fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Create(&fact)

	return c.Status(200).JSON(fact)
}
