package templates

import (
	"embed"
	"html/template"
)

//go:embed *.html
var content embed.FS

func Template(filename string) *template.Template {
	tmpl, _ := template.ParseFS(content, filename+".html")
	return tmpl
}
