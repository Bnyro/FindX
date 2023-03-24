package handlers

import (
	"net/http"
	"strconv"

	"github.com/bnyro/findx/utilities"
	"github.com/bnyro/findx/web"
)

func Api(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	searchType := r.URL.Query().Get("type")
	pageQuery := r.URL.Query().Get("page")
	page := 1
	if !utilities.IsBlank(pageQuery) {
		page, _ = strconv.Atoi(pageQuery)
	}

	if utilities.IsBlank(query) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := GenerateSearchMap(query, searchType, page)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	web.WriteJson(w, response)
}
