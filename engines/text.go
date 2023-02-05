package engines

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/bnyrogo/entities"
	"github.com/valyala/fasthttp"
)

func FetchText(query string, page int) ([]entities.Result, error) {
	var results []entities.Result

	req := fasthttp.AcquireRequest()

	uri := fmt.Sprintf("https://www.google.com/search?q=%s&start=%d", query, (page - 1) * 10)
	req.SetRequestURI(uri)
	req.Header.SetCookie("CONSENT", "YES+")
	resp := fasthttp.AcquireResponse()
    defer fasthttp.ReleaseResponse(resp)

	err := fasthttp.Do(req, resp)

	if err != nil {
		return results, err
	}

	contentEncoding := resp.Header.Peek("Content-Encoding")
    var body []byte
    if bytes.EqualFold(contentEncoding, []byte("gzip")) {
        body, _ = resp.BodyGunzip()
    } else {
        body = resp.Body()
    }

	reader := strings.NewReader(string(body))
    doc, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		fmt.Println(err.Error())
	}

	doc.Find(".Gx5Zad").Each(func(i int, s *goquery.Selection) {
		result := entities.Result{}
		href := s.Children().First().Children().First()
		url, _ := href.Attr("href")

		if !strings.Contains(url, "http") { return }

		var re = regexp.MustCompile("&sa=.*")
		result.Href = re.ReplaceAllString(url[7:], "")

		title := href.Children().First().Children().First().Children().First().Children().First()
		result.Title = strings.TrimSpace(title.Text())

		description := s.Children().Last().Children().First().Children().First()
		result.Description = strings.TrimSpace(description.Text())

		results = append(results, result)
	})

	return results, nil
}