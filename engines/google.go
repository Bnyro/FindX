package engines

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/bnyro/findx/entities"
	"github.com/bnyro/findx/utilities"
	"github.com/bnyro/findx/web"
)

func FetchText(query string, page int) ([]entities.Result, error) {
	var results []entities.Result

	uri := fmt.Sprintf("https://www.google.com/search?q=%s&start=%d&ie=utf8&oe=utf8", query, (page-1)*10)
	doc, err := web.RequestHtml(uri)

	if err != nil {
		return results, err
	}

	doc.Find(".Gx5Zad").Each(func(i int, s *goquery.Selection) {
		result := entities.Result{}
		href := s.Children().First().Children().First()
		link, _ := href.Attr("href")

		if !strings.Contains(link, "http") {
			return
		}

		var re = regexp.MustCompile("&sa=.*")
		url, _ := url.QueryUnescape(re.ReplaceAllString(link[7:], ""))
		result.Url = utilities.Redirect(url)

		short := href.Children().Last().Children().Last()
		result.Short = short.Text()

		title := href.Children().First().Children().First().Children().First().Children().First()
		result.Title = strings.TrimSpace(title.Text())

		description := s.Children().Last().Children().First().Children().First()
		result.Description = strings.TrimSpace(description.Text())

		results = append(results, result)
	})

	return results, nil
}
