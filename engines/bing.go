package engines

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/bnyro/findx/entities"
	"github.com/bnyro/findx/utilities"
	"github.com/bnyro/findx/web"
)

func FetchTextFallback(query string, page int) ([]entities.Result, error) {
	var results []entities.Result

	uri := fmt.Sprintf("https://www.bing.com/search?q=%s&first=%d", query, (page-1)*10+1)
	doc, err := web.RequestHtml(uri)

	if err != nil {
		return results, err
	}

	doc.Find(".b_algo").Each(func(i int, s *goquery.Selection) {
		result := entities.Result{}

		result.Title = s.Find(".b_title a").First().Text()
		if utilities.IsBlank(result.Title) {
			result.Title = s.Find("a").First().Text()
		}

		result.Url, _ = s.Find("a").First().Attr("href")
		result.Short = utilities.HumanizeUrl(result.Url)
		result.Description = s.Find(".b_caption p").First().Text()

		if utilities.IsBlank(result.Description) {
			result.Description = s.Find(".b_algoSlug").Text()
		}
		if utilities.IsBlank(result.Description) {
			result.Description = s.Find(".b_paractl").Text()
		}

		if !utilities.IsBlank(result.Description) {
			results = append(results, result)
		}
	})

	return results, nil
}
