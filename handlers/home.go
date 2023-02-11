package handlers

import (
	"net/http"

	"github.com/bnyro/findx/templates"
)

func Home(w http.ResponseWriter, r *http.Request) {
	templates.Template("home").Execute(w, nil)
}
