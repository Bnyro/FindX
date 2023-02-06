package handlers

import (
	"net/url"
	"strconv"

	"github.com/bnyrogo/engines"
	"github.com/bnyrogo/entities"
	"github.com/gofiber/fiber/v2"
)

func Search(c *fiber.Ctx) error {
	query := url.QueryEscape(c.Query("q", ""))
	sort := c.Query("type")
	page, err := strconv.Atoi(c.Query("page", "1"))

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	var results []entities.Result
	var images []entities.Image
	var videos []entities.Video
	switch sort {
	case "image": images, err = engines.FetchImage(query, page)
	default: results, err = engines.FetchText(query, page)
	}

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Render("results", fiber.Map {
		"Query": query,
		"Page": page,
		"Results": results,
		"Images": images,
		"Videos": videos,
	})
}