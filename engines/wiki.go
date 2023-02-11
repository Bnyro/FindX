package engines

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/bnyro/findx/entities"
	"github.com/bnyro/findx/utilities"
	"github.com/bnyro/findx/web"
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
		description := utilities.TakeN(entry["extract"].(string), 350)

		if description == "" || strings.Contains(description, "may refer to:") {
			return result, errors.New("Not found")
		}

		result.Description = description
		switch entry["thumbnail"].(type) {
		case map[string]interface{}:
			thumbnail := entry["thumbnail"].(map[string]interface{})["source"].(string)
			result.Thumbnail = utilities.RewriteProxied(thumbnail)
		default:
		}
		break
	}

	return result, nil
}
