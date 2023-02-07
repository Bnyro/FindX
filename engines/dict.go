package engines

import (
	"errors"
	"fmt"

	"github.com/bnyrogo/entities"
	"github.com/bnyrogo/web"
)

func FetchDictionary(query string) (entities.Dict, error) {
	var dict entities.Dict

	uri := fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", query)
	var data []map[string]interface{}

	err := web.RequestJson(uri, &data)

	if err != nil {
		return dict ,err
	}

	if len(data) == 0 {
		return dict, errors.New("Not found")
	}

	meanings := data[0]["meanings"].([]interface{})

	outer: for _, entry := range meanings {
		switch entry.(type) {
		case map[string]interface{}: {
			meaning := entry.(map[string]interface{})
			dict.PartOfSpeech = meaning["partOfSpeech"].(string)
			for _, definition := range meaning["definitions"].([]interface{}) {
				def := definition.(map[string]interface{})
				dict.Definition = def["definition"].(string)
				val, ok := def["example"]
				if ok {
					dict.Example = val.(string)
				}
				break outer
			}
			}
		}
	}

	return dict, nil
}