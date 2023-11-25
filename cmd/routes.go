package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lukke-dev/api-golang/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)
	app.Get("/fact/:id", handlers.ShowFact)
	app.Post("/fact", handlers.CreateFact)
	app.Delete("/fact/:id", handlers.DeleteFact)
}