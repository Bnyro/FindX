package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/bnyro/findx/config"
	"github.com/bnyro/findx/handlers"
	"github.com/bnyro/findx/web"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

var Main string

func main() {
	config.Init()

	app := chi.NewRouter()

	app.Use(cors.AllowAll().Handler)

	app.Get("/", handlers.Home)
	app.Get("/search", handlers.Search)
	app.Get("/api", handlers.Api)
	app.Get("/ac", handlers.Suggest)

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "static"))
	web.FileServer(app, "/static", filesDir)
	app.Get("/proxy", handlers.Proxy)
	app.Get("/config", handlers.Config)
	app.Get("/opensearch.xml", handlers.Opensearch)

	fmt.Printf("Listening on: http://localhost:%s\n", *config.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", *config.Port), app)
}
