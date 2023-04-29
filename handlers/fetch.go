package handlers

import (
	"html/template"
	"net/url"
	"sync"
	"time"

	"github.com/bnyro/findx/engines"
	"github.com/bnyro/findx/entities"
	"github.com/bnyro/findx/utilities"
	"github.com/bnyro/findx/web"
)

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
	var news []entities.News
	var code []entities.Stack
	var videos []entities.Video
	switch searchType {
	case "image":
		images, err = engines.FetchImage(escapedQuery, page)
		if err != nil || len(images) == 0 {
			images, err = engines.FetchImageFallback(escapedQuery, page)
		}
	case "news":
		news, err = engines.FetchNews(escapedQuery)
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
			if err != nil || len(results) == 0 {
				results, err = engines.FetchTextFallback(escapedQuery, page)
			}
			// wait at most one additional second for the additional results
			utilities.WaitTimeout(&wg, 1*time.Second)
		}
	}

	if err != nil {
		return nil, err
	}

	timeTaken := time.Since(start)

	return web.Map{
		"query":     query,
		"type":      searchType,
		"page":      page,
		"timeTaken": timeTaken.String(),
		"providers": providers,
		"wiki":      wiki,
		"dict":      dict,
		"weather":   template.HTML(weather),
		"results":   results,
		"images":    images,
		"news":      news,
		"code":      code,
		"videos":    videos,
	}, nil
}
