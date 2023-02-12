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
		return "", errors.New("Query must contain alphabetic chars only")
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

	area := data["nearest_area"].([]interface{})[0]
	areaName := area.(map[string]interface{})["areaName"].([]interface{})[0]
	location := strings.ToLower(areaName.(map[string]interface{})["value"].(string))

	region := area.(map[string]interface{})["region"].([]interface{})[0]
	regionName := strings.ToLower(region.(map[string]interface{})["value"].(string))

	country := area.(map[string]interface{})["country"].([]interface{})[0]
	countryName := strings.ToLower(country.(map[string]interface{})["value"].(string))

	lowerQuery := strings.ToLower(query)
	if location != lowerQuery && regionName != lowerQuery && countryName != lowerQuery {
		return "", errors.New("Random result")
	}

	weather, _ := doc.Find("pre").First().Html()
	return weather, nil
}
