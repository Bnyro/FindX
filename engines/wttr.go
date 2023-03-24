package engines

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/bnyro/findx/utilities"
	"github.com/bnyro/findx/web"
)

func FetchWeather(query string) (string, error) {
	if !utilities.IsAlphabetic(query) {
		return "", errors.New("query must contain alphabetic chars only")
	}

	uri := fmt.Sprintf("https://wttr.in/%s?T?0", query)
	confirmUri := fmt.Sprintf("https://wttr.in/%s?format=j1", query)

	var err error
	var data web.Map
	var doc *goquery.Document

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		web.RequestJson(confirmUri, &data)
	}()
	go func() {
		defer wg.Done()
		doc, err = web.RequestHtml(uri)
	}()
	wg.Wait()

	if err != nil {
		return "", err
	}

	areaName := extractValue(data, "areaName")
	regionName := extractValue(data, "region")
	countryName := extractValue(data, "country")

	lowerQuery := strings.ToLower(query)
	if areaName != lowerQuery && regionName != lowerQuery && countryName != lowerQuery {
		return "", errors.New("invalid result")
	}

	weather, _ := doc.Find("pre").First().Html()
	return weather, nil
}

func extractValue(data web.Map, key string) string {
	if data["nearest_area"] == nil {
		return ""
	}
	areas := data["nearest_area"].([]interface{})
	if len(areas) == 0 {
		return ""
	}
	area := areas[0]
	value := area.(map[string]interface{})[key].([]interface{})[0]
	return strings.ToLower(value.(map[string]interface{})["value"].(string))
}
