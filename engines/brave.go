package engines

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/bnyro/findx/entities"
	"github.com/bnyro/findx/utilities"
	"github.com/bnyro/findx/web"
)

func FetchNews(query string) ([]entities.News, error) {
	var news []entities.News

	uri := fmt.Sprintf("https://search.brave.com/news?q=%s", query)

	doc, err := web.RequestHtml(uri)

	if err != nil {
		return news, err
	}

	doc.Find(".snippet").Each(func(i int, s *goquery.Selection) {
		entry := entities.News{}
		entry.Title = s.Find(".snippet-title").Text()
		entry.Description = s.Find(".snippet-description").Text()
		url, _ := s.Find(".result-header").Attr("href")
		entry.Url = utilities.Redirect(url)
		entry.Source = s.Find(".netloc").Text()
		entry.UploadDate = s.Find(".snippet-url").Children().Last().Text()
		thumbnail, _ := s.Find(".thumb").Attr("src")
		if !utilities.IsBlank(thumbnail) {
			entry.Thumbnail = utilities.RewriteProxied(thumbnail)
		}
		news = append(news, entry)
	})

	return news, nil
}
