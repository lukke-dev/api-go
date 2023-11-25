package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lukke-dev/api-golang/database"
	"github.com/lukke-dev/api-golang/models"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}

	database.DB.Db.Find(&facts)

	return c.Status(fiber.StatusOK).JSON(facts)
}

func ShowFact(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fiber.ErrBadRequest.Message,
		})
	}

	fact := models.Fact{ID: id}
	result := database.DB.Db.Find(&fact)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fiber.ErrNotFound.Message,
		})
	}

	return c.Status(fiber.StatusFound).JSON(fact)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)

	err := c.BodyParser(fact)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)
	return c.Status(fiber.StatusOK).JSON(fact)
}

func DeleteFact(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fiber.ErrBadRequest.Message,
		})
	}

	fact := models.Fact{ID: id}
	result := database.DB.Db.Find(&fact)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fiber.ErrNotFound.Message,
		})
	}

	database.DB.Db.Delete(&fact)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Fact deleted Successfully",
	})
}