package engines

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/bnyrogo/entities"
	"github.com/bnyrogo/utilities"
	"github.com/bnyrogo/web"
)

const wikiUrl = "https://en.wikipedia.org"

func FetchWiki(query string) (entities.Wiki, error) {
	filter := url.QueryEscape("extracts|pageimages")
	q := url.QueryEscape(query)
	uri := fmt.Sprintf("%s/w/api.php?format=json&action=query&prop=%s&exintro&explaintext&redirects=1&pithumbsize=500&titles=%s", wikiUrl, filter, q)
	result := entities.Wiki{
		Url: uri,
	}

	var data map[string]interface{}
	err := web.RequestJson(uri, &data)

	if err != nil {
		return result, err
	}

	pages := data["query"].(map[string]interface{})["pages"].(map[string]interface{})
	for key, value := range pages {
		if key == "-1" {
			break
		}
		entry := value.(map[string]interface{})
		result.Description = utilities.TakeN(entry["extract"].(string), 350)
		switch entry["thumbnail"].(type) {
		case map[string]interface{}:
			result.Thumbnail = entry["thumbnail"].(map[string]interface{})["source"].(string)
		default:
		}
	}

	if result.Description == "" {
		return result, errors.New("Not found")
	}

	return result, nil
}
