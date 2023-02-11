package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bnyro/findx/utilities"
)

type Map map[string]interface{}

func WriteJsonStatus(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	body, _ := json.Marshal(data)
	w.Write(body)
}

func WriteJson(w http.ResponseWriter, data interface{}) {
	WriteJsonStatus(w, data, http.StatusOK)
}

func Host(r *http.Request) string {
	host := r.Host
	forwardedHost := r.Header.Get("X-Forwarded-Host")
	if !utilities.IsBlank(forwardedHost) {
		return forwardedHost
	}
	proto := r.Header.Get("X-Forwarded-Proto")
	if utilities.IsBlank(proto) {
		proto = "http"
	}
	return fmt.Sprintf("%s://%s", proto, host)
}
