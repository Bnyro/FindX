package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bnyro/findx/entities"
	"github.com/bnyro/findx/templates"
	"github.com/bnyro/findx/utilities"
	"github.com/bnyro/findx/web"
)

// List of all available search types
var providers = []entities.Type{
	{Query: "text", Name: "General"},
	{Query: "image", Name: "Image"},
	{Query: "news", Name: "News"},
	{Query: "code", Name: "Code"},
	{Query: "video", Name: "Video"},
	{Query: "music", Name: "Music"},
}

func Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	searchType := r.URL.Query().Get("type")
	pageQuery := r.URL.Query().Get("page")
	page := 1
	if !utilities.IsBlank(pageQuery) {
		page, _ = strconv.Atoi(pageQuery)
	}

	if utilities.IsBlank(query) {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	response, err := GenerateSearchMap(query, searchType, page)
	if err != nil {
		fmt.Println(err)
		response = web.Map{
			"error": err.Error(),
			"query": query,
			"page":  page,
			"type":  searchType,
		}
	}

	err = templates.Template("results").Execute(w, response)

	if err != nil {
		fmt.Println(err.Error())
	}
}
