package handlers

import (
	"net/http"
	"strings"

	_ "embed"

	"github.com/bnyro/findx/config"
	"github.com/bnyro/findx/web"
)

//go:embed opensearch.xml
var opensearchBody string

func Config(w http.ResponseWriter, r *http.Request) {
	web.WriteJson(w, web.Map{
		"proxy":     *config.Proxy,
		"redirects": config.Redirects,
	})
}

func Opensearch(w http.ResponseWriter, r *http.Request) {
	descr := strings.Replace(opensearchBody, "{{baseUrl}}", r.URL.Host, -1)
	w.Write([]byte(descr))
}
