package engines

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/bnyro/findx/entities"
	"github.com/bnyro/findx/utilities"
	"github.com/bnyro/findx/web"
)

func FetchImageFallback(query string, page int) ([]entities.Image, error) {
	var results []entities.Image

	uri := fmt.Sprintf("https://images.search.yahoo.com/search/images?p=%s&ei=UTF-8&b=%d", query, (page-1)*50)

	doc, err := web.RequestHtml(uri)

	if err != nil {
		return results, err
	}

	doc.Find("#sres li .img").Each(func(i int, s *goquery.Selection) {
		var result entities.Image
		imgSrc := utilities.RewriteProxied(strings.Split(s.Text(), "'")[1])
		result.Thumbnail = imgSrc
		result.Media = imgSrc
		result.Title, _ = s.Attr("aria-label")
		urlComponent, _ := s.Attr("href")
		fullUrl, _ := url.QueryUnescape(strings.SplitN(urlComponent, "RU=", 2)[1])
		result.Url = strings.Split(fullUrl, "/RK=")[0]
		results = append(results, result)
	})

	return results, nil
}
