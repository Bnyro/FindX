package engines

import (
	"fmt"

	"github.com/bnyrogo/web"
)

func GetSuggestions(query string) []string {
	var results []string

	uri := fmt.Sprintf("https://duckduckgo.com/ac/?q=%s", query)

	var resp []map[string]interface{}

	err := web.RequestJson(uri, &resp)

	if err != nil {
		return results
	}

	for _, res := range resp {
		results = append(results, res["phrase"].(string))
	}

	return results
}
