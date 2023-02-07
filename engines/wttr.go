package engines

import (
	"fmt"

	"github.com/bnyrogo/web"
)

func FetchWeather(query string) (string, error) {
	uri := fmt.Sprintf("https://wttr.in/%s?H?0", query)
	resp, err := web.Request(uri)

	if err != nil {
		return "", err
	}

	return string(resp), nil
}
