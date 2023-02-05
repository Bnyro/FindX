package handlers

import (
	"net/url"
	"strconv"

	"github.com/bnyrogo/engines"
	"github.com/gofiber/fiber/v2"
)

func Search(c *fiber.Ctx) error {
	query := url.QueryEscape(c.Query("q", ""))
	page, _ := strconv.Atoi(c.Query("page", "1"))
	results, err := engines.FetchText(query, page)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Render("results", fiber.Map {
		"Query": query,
		"Page": page,
		"Results": results,
	})
}