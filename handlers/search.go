package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/bnyro/findx/engines"
	"github.com/bnyro/findx/entities"
	"github.com/bnyro/findx/templates"
	"github.com/bnyro/findx/utilities"
	"github.com/bnyro/findx/web"
)

// List of all available search types
var providers = []entities.Type{
	{Query: "text", Name: "General"},
	{Query: "image", Name: "Image"},
	{Query: "news", Name: "News"},
	{Query: "code", Name: "Code"},
	{Query: "video", Name: "Video"},
	{Query: "music", Name: "Music"},
}

func Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	searchType := r.URL.Query().Get("type")
	pageQuery := r.URL.Query().Get("page")
	page := 1
	if !utilities.IsBlank(pageQuery) {
		page, _ = strconv.Atoi(pageQuery)
	}

	if utilities.IsBlank(query) {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	response, err := GenerateSearchMap(query, searchType, page)
	if err != nil {
		fmt.Println(err)
		response = web.Map{
			"error": err.Error(),
			"query": query,
			"page":  page,
			"type":  searchType,
		}
	}

	templates.Template("results").Execute(w, response)
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
	var news []entities.News
	var code []entities.Stack
	var videos []entities.Video
	switch searchType {
	case "image":
		images, err = engines.FetchImage(escapedQuery, page)
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
