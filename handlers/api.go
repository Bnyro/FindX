package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Api(c *fiber.Ctx) error {
	response, err := GenerateSearchMap(c)

	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(response)
}
