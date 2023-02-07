package engines

import (
	"encoding/json"
	"fmt"
	"html"

	"github.com/bnyrogo/entities"
	"github.com/bnyrogo/utilities"
	"github.com/bnyrogo/web"
)

const stackUrl = "https://api.stackexchange.com"

func FetchCode(query string, page int) ([]entities.Stack, error) {
	var stacks []entities.Stack
	uri := fmt.Sprintf("%s/search/advanced?order=desc&sort=relevance&site=stackoverflow&q=%s&page=%d", stackUrl, query, page)

	var data map[string]interface{}
	err := web.RequestJson(uri, &data)

	if err != nil {
		return stacks, err
	}

	items, _ := json.Marshal(data["items"])

	err = json.Unmarshal([]byte(items), &stacks)

	if err != nil {
		return stacks, err
	}

	for i := range stacks {
		stacks[i].Title = html.UnescapeString(stacks[i].Title)
		stacks[i].ScoreStr = utilities.HumanReadable(uint64(stacks[i].Score))
		stacks[i].ViewCountStr = utilities.HumanReadable(stacks[i].ViewCount)
		stacks[i].CreationDateStr = utilities.ParseUnixTime(stacks[i].CreationDate * 1000)
	}

	return stacks, nil
}
