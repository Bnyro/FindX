package engines

import (
	"encoding/json"
	"fmt"

	"github.com/bnyrogo/entities"
	"github.com/bnyrogo/utilities"
	"github.com/bnyrogo/web"
)

const pipedApiUrl = "https://pipedapi-libre.kavin.rocks"
const pipedUrl = "https://piped.kavin.rocks"

func FetchVideo(query string) ([]entities.Video, error) {
	var videos []entities.Video

	uri := fmt.Sprintf("%s/search?q=%s&filter=video", pipedApiUrl, query)

	var data map[string]interface{}
	err := web.RequestJson(uri, &data)

	if err != nil {
		return videos, err
	}

	jsonVids, _ := json.Marshal(data["items"])
	json.Unmarshal(jsonVids, &videos)

	for i := range videos {
		videos[i].Url = pipedUrl + videos[i].Url
		videos[i].DurationString = utilities.FmtDuration(videos[i].Duration)
		videos[i].UploadDate = utilities.ParseUnixTime(videos[i].Uploaded)
		videos[i].ViewsString = utilities.HumanReadable(videos[i].Views)
	}

	return videos, nil
}