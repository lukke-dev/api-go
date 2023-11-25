package main

import (
  "github.com/gofiber/fiber/v2"
  "github.com/lukke-dev/api-golang/database"
)

func main() {
  database.ConnectDb()
  app := fiber.New()

  app.Get("/", func (c *fiber.Ctx) error {
      return c.SendString("Hello, World! hot reload")
  })

  app.Listen(":3000")
}