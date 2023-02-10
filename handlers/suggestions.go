package handlers

import (
	"github.com/bnyro/findx/engines"
	"github.com/gofiber/fiber/v2"
)

func Suggest(c *fiber.Ctx) error {
	query := c.Query("q")
	results := engines.GetSuggestions(query)
	return c.JSON([]any{
		query, results,
	})
}
