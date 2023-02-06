package handlers

import (
	"net/url"
	"strconv"
	"time"

	"github.com/bnyrogo/engines"
	"github.com/bnyrogo/entities"
	"github.com/gofiber/fiber/v2"
)

func Api(c *fiber.Ctx) error {
	start := time.Now()

	query := c.Query("q", "")
	escapedQuery := url.QueryEscape(query)

	searchType := c.Query("type")
	page, err := strconv.Atoi(c.Query("page", "1"))

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	var results []entities.Result
	var images []entities.Image
	var videos []entities.Video
	switch searchType {
	case "image": images, err = engines.FetchImage(escapedQuery, page)
	case "video": videos, err = engines.FetchVideo(escapedQuery)
	case "music": videos, err = engines.FetchMusic(escapedQuery)
	default: {
		results, err = engines.FetchText(escapedQuery, page)
		searchType = "text"
	}
	}

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	timeTaken := time.Since(start)

	return c.JSON(fiber.Map {
		"query": query,
		"type": searchType,
		"page": page,
		"timeTaken": timeTaken,
		"results": results,
		"images": images,
		"videos": videos,
	})
}