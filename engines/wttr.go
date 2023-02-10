package engines

import (
	"fmt"

	"github.com/bnyro/findx/web"
)

func FetchWeather(query string) (string, error) {
	uri := fmt.Sprintf("https://wttr.in/%s?T?0", query)
	doc, err := web.RequestHtml(uri)

	if err != nil {
		return "", err
	}

	weather, _ := doc.Find("pre").First().Html()
	return weather, nil
}
