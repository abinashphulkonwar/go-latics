package routes

import (
	"github.com/abinashphulkonwar/go-latics/handlers"
	"github.com/gofiber/fiber/v2"
)

func AddRoutes(ctx fiber.Router) {
	ctx.Post("/views", handlers.ViewsHandler)
}
