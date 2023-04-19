package main

import (
	"github.com/abinashphulkonwar/go-latics/db"
	"github.com/abinashphulkonwar/go-latics/routes"
	"github.com/gofiber/fiber/v2"
)

func App() *fiber.App {
	_, err := db.InitClient()
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	app.Route("/add", routes.AddRoutes)

	app.Use(func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"Title": "data",
		})
	})

	return app
}
