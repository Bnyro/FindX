package handlers

import (
	"github.com/bnyrogo/utilities"
	"github.com/gofiber/fiber/v2"
)

func Api(c *fiber.Ctx) error {
	if utilities.IsBlank(c.Query("q")) {
		return c.SendStatus(400)
	}

	response, err := GenerateSearchMap(c)

	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(response)
}
