package engines

import (
	"encoding/json"
	"fmt"

	"github.com/bnyrogo/entities"
	"github.com/bnyrogo/web"
)

const stackUrl = "https://api.stackexchange.com"

func FetchCode(query string, page int) ([]entities.Stack, error) {
	var stacks []entities.Stack
	uri := fmt.Sprintf("%s/search/advanced?order=desc&sort=relevance&site=stackoverflow&q=%s", stackUrl, query)

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

	return stacks, nil
}
