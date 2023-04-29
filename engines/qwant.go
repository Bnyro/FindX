package engines

import (
	"errors"
	"fmt"

	"github.com/bnyro/findx/entities"
	"github.com/bnyro/findx/utilities"
	"github.com/bnyro/findx/web"
)

type QwantImageResponse struct {
	Status string `json:"status"`
	Data   struct {
		Query struct {
			Locale       string `json:"locale"`
			Query        string `json:"query"`
			Offset       int    `json:"offset"`
			QueryContext struct {
				OriginalQuery string `json:"originalQuery"`
			} `json:"queryContext"`
		} `json:"query"`
		Result struct {
			Total int `json:"total"`
			Items []struct {
				Title         string `json:"title"`
				Media         string `json:"media"`
				Thumbnail     string `json:"thumbnail"`
				ThumbWidth    int    `json:"thumb_width"`
				ThumbHeight   int    `json:"thumb_height"`
				Width         int    `json:"width"`
				Height        int    `json:"height"`
				URL           string `json:"url"`
				ID            string `json:"_id"`
				MediaFullsize string `json:"media_fullsize"`
				MediaPreview  string `json:"media_preview"`
				Size          string `json:"size"`
				ThumbType     string `json:"thumb_type"`
			} `json:"items"`
			Instrumentation struct {
				PingURLBase     string `json:"pingUrlBase"`
				PageLoadPingURL string `json:"pageLoadPingUrl"`
			} `json:"instrumentation"`
			Filters struct {
				Size struct {
					Label    string `json:"label"`
					Name     string `json:"name"`
					Type     string `json:"type"`
					Selected string `json:"selected"`
					Values   []struct {
						Value     string `json:"value"`
						Label     string `json:"label"`
						Translate bool   `json:"translate"`
					} `json:"values"`
				} `json:"size"`
				License struct {
					Label    string `json:"label"`
					Name     string `json:"name"`
					Type     string `json:"type"`
					Selected string `json:"selected"`
					Values   []struct {
						Value     string `json:"value"`
						Label     string `json:"label"`
						Translate bool   `json:"translate"`
					} `json:"values"`
				} `json:"license"`
				Freshness struct {
					Label    string `json:"label"`
					Name     string `json:"name"`
					Type     string `json:"type"`
					Selected string `json:"selected"`
					Values   []struct {
						Value     string `json:"value"`
						Label     string `json:"label"`
						Translate bool   `json:"translate"`
					} `json:"values"`
				} `json:"freshness"`
				Color struct {
					Label    string `json:"label"`
					Name     string `json:"name"`
					Type     string `json:"type"`
					Selected string `json:"selected"`
					Values   []struct {
						Value     string `json:"value"`
						Label     string `json:"label"`
						Translate bool   `json:"translate"`
					} `json:"values"`
				} `json:"color"`
				Imagetype struct {
					Label    string `json:"label"`
					Name     string `json:"name"`
					Type     string `json:"type"`
					Selected string `json:"selected"`
					Values   []struct {
						Value     string `json:"value"`
						Label     string `json:"label"`
						Translate bool   `json:"translate"`
					} `json:"values"`
				} `json:"imagetype"`
			} `json:"filters"`
			LastPage bool `json:"lastPage"`
		} `json:"result"`
	} `json:"data"`
}

const resultsPerPage = 25

func FetchImage(query string, page int) ([]entities.Image, error) {
	var images []entities.Image
	var data QwantImageResponse
	offset := (page - 1) * resultsPerPage

	if offset+resultsPerPage >= 250 {
		return images, errors.New("count + offset must be smaller than 250")
	}

	uri := fmt.Sprintf("https://api.qwant.com/v3/search/images?q=%s&offset=%d&locale=en_gb&count=%d", query, offset, resultsPerPage)
	err := web.RequestJson(uri, &data)

	if err != nil {
		return images, err
	}

	for _, result := range data.Data.Result.Items {
		image := entities.Image{
			Title:     result.Title,
			Url:       utilities.Redirect(result.URL),
			Thumbnail: utilities.RewriteProxied(result.Thumbnail),
			Media:     result.Media,
		}
		images = append(images, image)
	}

	return images, nil
}
