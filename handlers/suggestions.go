package handlers

import (
	"github.com/bnyrogo/engines"
	"github.com/gofiber/fiber/v2"
)

func Suggest(c *fiber.Ctx) error {
	query := c.Query("q")
	results := engines.GetSuggestions(query)
	return c.JSON(results)
}
