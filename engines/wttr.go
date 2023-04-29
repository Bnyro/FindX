package engines

import (
	"errors"
	"fmt"

	"github.com/bnyro/findx/utilities"
	"github.com/bnyro/findx/web"
)

func FetchWeather(query string) (string, error) {
	if !utilities.IsAlphabetic(query) {
		return "", errors.New("query must contain alphabetic chars only")
	}

	uri := fmt.Sprintf("https://wttr.in/~%s?0", query)

	var err error

	doc, err := web.RequestHtml(uri)

	if err != nil {
		return "", err
	}

	return doc.Find("pre").First().Html()
}
