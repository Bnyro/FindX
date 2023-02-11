package web

import (
	"net/http"

	"github.com/goccy/go-json"
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
