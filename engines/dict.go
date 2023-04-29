package engines

import (
	"errors"
	"fmt"

	"github.com/bnyro/findx/entities"
	"github.com/bnyro/findx/web"
)

type DictResponse []struct {
	Word      string `json:"word"`
	Phonetics []struct {
		Audio     string `json:"audio"`
		SourceURL string `json:"sourceUrl,omitempty"`
		License   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"license,omitempty"`
		Text string `json:"text,omitempty"`
	} `json:"phonetics"`
	Meanings []struct {
		PartOfSpeech string `json:"partOfSpeech"`
		Definitions  []struct {
			Definition string `json:"definition"`
			Synonyms   []any  `json:"synonyms"`
			Antonyms   []any  `json:"antonyms"`
			Example    string `json:"example"`
		} `json:"definitions"`
		Synonyms []string `json:"synonyms"`
		Antonyms []any    `json:"antonyms"`
	} `json:"meanings"`
	License struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"license"`
	SourceUrls []string `json:"sourceUrls"`
}

func FetchDictionary(query string) (entities.Dict, error) {
	var dict entities.Dict

	uri := fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", query)
	var data DictResponse

	err := web.RequestJson(uri, &data)

	if err != nil {
		return dict, err
	}

	if len(data) == 0 {
		return dict, errors.New("not found")
	}

outer:
	for _, entry := range data[0].Meanings {
		dict.PartOfSpeech = entry.PartOfSpeech
		for _, definition := range entry.Definitions {
			dict.Definition = definition.Definition
			dict.Example = definition.Example
			break outer
		}
	}

	return dict, nil
}
