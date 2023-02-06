package engines

import (
	"errors"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/bnyrogo/entities"
	"github.com/bnyrogo/web"
)

const wikilessUrl = "https://wiki.adminforge.de"
func FetchWiki(query string) (entities.Wiki, error) {
	uri := fmt.Sprintf("%s/wiki/%s?lang=en", wikilessUrl, query)
	result := entities.Wiki{
		Url: uri,
	}

	doc, err := web.RequestHtml(uri)

	if err != nil {
		return result, err
	}

	
	thumbnail, exists := doc.Find(".thumbimage").First().Attr("src")
	if exists {
		result.Thumbnail = wikilessUrl + thumbnail
	}
	
	doc.Find(".mw-parser-output").Children().Each(func(i int, s *goquery.Selection) {
		if result.Description != "" || strings.TrimSpace(s.Text()) == "" { return }
		if s.Is("p") && !s.HasClass("mv-empty-elt") {
			result.Description = s.Text()[:500] + " ..."
		}
	})

	if (result.Description == "") { return result, errors.New("Not found") }

	return result, nil
}