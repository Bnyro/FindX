package handlers

import (
	"net/http"

	"github.com/bnyro/findx/utilities"
	"github.com/bnyro/findx/web"
)

func Proxy(w http.ResponseWriter, r *http.Request) {
	uri := r.URL.Query().Get("url")

	if utilities.IsBlank(uri) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, contentType, err := web.Request(uri)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", string(contentType))
	w.Write(body)
}
