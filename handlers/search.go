package handlers

import (
	"fmt"
	"html/template"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/bnyrogo/engines"
	"github.com/bnyrogo/entities"
	"github.com/bnyrogo/utilities"
	"github.com/gofiber/fiber/v2"
)

func Search(c *fiber.Ctx) error {
	query := c.Query("q", "")
	searchType := c.Query("type")
	page, err := strconv.Atoi(c.Query("page", "1"))

	if err != nil {
		return c.SendStatus(403)
	}

	if utilities.IsBlank(query) {
		return c.Redirect("/")
	}

	response, err := GenerateSearchMap(query, searchType, page)
	if err != nil {
		return c.Render("results", fiber.Map{
			"error": err.Error(),
			"query": query,
			"page":  page,
			"type":  searchType,
		})
	}

	return c.Render("results", response)
}

func GenerateSearchMap(query string, searchType string, page int) (map[string]interface{}, error) {
	start := time.Now()
	var err error

	if page < 1 {
		page = 1
	}

	escapedQuery := url.QueryEscape(query)

	var wiki entities.Wiki
	var dict entities.Dict
	var weather string
	var results []entities.Result
	var images []entities.Image
	var code []entities.Stack
	var videos []entities.Video
	switch searchType {
	case "image":
		images, err = engines.FetchImage(escapedQuery, page)
	case "code":
		code, err = engines.FetchCode(escapedQuery, page)
	case "video":
		videos, err = engines.FetchVideo(escapedQuery)
	case "music":
		videos, err = engines.FetchMusic(escapedQuery)
	default:
		{
			searchType = "text"
			var wg sync.WaitGroup

			// only show meta results on first page
			if page == 1 {
				wg.Add(3)
				go func() {
					defer wg.Done()
					wiki, _ = engines.FetchWiki(query)
				}()
				go func() {
					defer wg.Done()
					dict, _ = engines.FetchDictionary(query)
				}()
				go func() {
					defer wg.Done()
					weather, _ = engines.FetchWeather(query)
				}()
			}

			results, err = engines.FetchText(escapedQuery, page)
			// wait at most one additional second for the additional results
			utilities.WaitTimeout(&wg, 1*time.Second)
		}
	}

	if err != nil {
		return nil, err
	}

	timeTaken := time.Since(start)

	return fiber.Map{
		"query":     query,
		"type":      searchType,
		"page":      page,
		"timeTaken": fmt.Sprintf("%s", timeTaken),
		"wiki":      wiki,
		"dict":      dict,
		"weather":   template.HTML(weather),
		"results":   results,
		"images":    images,
		"code":      code,
		"videos":    videos,
	}, nil
}
