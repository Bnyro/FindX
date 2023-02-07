package handlers

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/bnyrogo/engines"
	"github.com/bnyrogo/entities"
	"github.com/gofiber/fiber/v2"
)

func Search(c *fiber.Ctx) error {
	response, err := GenerateSearchMap(c)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.Render("results", response)
}

func GenerateSearchMap(c *fiber.Ctx) (map[string]interface{}, error) {
	start := time.Now()

	query := c.Query("q", "")
	escapedQuery := url.QueryEscape(query)
	searchType := c.Query("type")
	page, err := strconv.Atoi(c.Query("page", "1"))

	if err != nil {
		return nil, err
	}

	if page < 1 { page = 1 }

	var wiki entities.Wiki
	var results []entities.Result
	var images []entities.Image
	var videos []entities.Video
	switch searchType {
	case "image": images, err = engines.FetchImage(escapedQuery, page)
	case "video": videos, err = engines.FetchVideo(escapedQuery)
	case "music": videos, err = engines.FetchMusic(escapedQuery)
	default: {
		wiki, _ = engines.FetchWiki(query)
		results, err = engines.FetchText(escapedQuery, page)
		searchType = "text"
	}
	}

	if err != nil {
		return nil, err
	}

	timeTaken := time.Since(start)

	return fiber.Map {
		"query": query,
		"type": searchType,
		"page": page,
		"prev": page - 1,
		"next": page + 1,
		"timeTaken": fmt.Sprintf("%s", timeTaken),
		"wiki": wiki,
		"results": results,
		"images": images,
		"videos": videos,
	}, nil
}