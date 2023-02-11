package engines

import (
	"encoding/json"
	"fmt"

	"github.com/bnyro/findx/entities"
	"github.com/bnyro/findx/utilities"
	"github.com/bnyro/findx/web"
)

const pipedApiUrl = "https://pipedapi-libre.kavin.rocks"
const ytUrl = "https://www.youtube.com"

func fetchSearch(query string, filter string) ([]entities.Video, error) {
	var videos []entities.Video

	uri := fmt.Sprintf("%s/search?q=%s&filter=%s", pipedApiUrl, query, filter)

	var data map[string]interface{}
	err := web.RequestJson(uri, &data)

	if err != nil {
		return videos, err
	}

	jsonVids, _ := json.Marshal(data["items"])
	json.Unmarshal(jsonVids, &videos)

	for i := range videos {
		videos[i].Url = utilities.Redirect(ytUrl + videos[i].Url)
		videos[i].DurationString = utilities.FormatDuration(videos[i].Duration)
		videos[i].UploadDate = utilities.FormatDate(videos[i].Uploaded)
		videos[i].ViewsString = utilities.FormatHumanReadable(int64(videos[i].Views))
		videos[i].Thumbnail = utilities.RewriteProxied(videos[i].Thumbnail)
	}

	return videos, nil
}

func FetchVideo(query string) ([]entities.Video, error) {
	return fetchSearch(query, "videos")
}

func FetchMusic(query string) ([]entities.Video, error) {
	return fetchSearch(query, "music_songs")
}
