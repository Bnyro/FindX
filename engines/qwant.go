package engines

import (
	"errors"
	"fmt"

	"github.com/bnyrogo/entities"
	"github.com/bnyrogo/web"
)

const resultsPerPage = 20

func FetchImage(query string, page int) ([]entities.Image, error) {
	var images []entities.Image
	var data map[string]interface{}
	offset := (page - 1) * resultsPerPage

	if offset + resultsPerPage >= 250 {
		return images, errors.New("Count + offset must be smaller than 250")
	}

	uri := fmt.Sprintf("https://api.qwant.com/v3/search/images?q=%s&offset=%d&locale=en_gb&count=%d", query, offset, resultsPerPage)
	err := web.RequestJson(uri, &data)

	if err != nil {
		return images, err
	}

	results := data["data"].(map[string]interface{})["result"].(map[string]interface{})["items"].([]interface{})
	for _, res := range results {
		result := res.(map[string]interface{})
		image := entities.Image{}
		image.Title = result["title"].(string)
		image.Url = result["url"].(string)
		image.Thumbnail = result["thumbnail"].(string)
		image.Media = result["media"].(string)
		images = append(images, image)
	}
	
	return images, nil
}