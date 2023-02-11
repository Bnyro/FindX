package handlers

import (
	"net/http"

	"github.com/bnyro/findx/engines"
	"github.com/bnyro/findx/web"
)

func Suggest(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	results := engines.GetSuggestions(query)
	web.WriteJson(w, []any{
		query, results,
	})
}
