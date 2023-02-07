package handlers

import (
	"strconv"

	"github.com/bnyrogo/utilities"
	"github.com/gofiber/fiber/v2"
)

func Api(c *fiber.Ctx) error {
	query := c.Query("q", "")
	searchType := c.Query("type")
	page, err := strconv.Atoi(c.Query("page", "1"))

	if utilities.IsBlank(query) {
		return c.SendStatus(400)
	}

	response, err := GenerateSearchMap(query, searchType, page)

	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(response)
}
