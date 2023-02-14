package config

import (
	"encoding/json"
	"flag"

	_ "embed"

	"github.com/bnyro/findx/entities"
)

var Port *string
var Proxy *bool
var Redirects []entities.Redirect

//go:embed redirects.json
var redirectsFile []byte

func Init() {
	json.Unmarshal(redirectsFile, &Redirects)

	Port = getSetting("port", "8080", "port to listen on")
	Proxy = getBool("proxy", false, "wether to proxy all images")

	flag.Parse()
}
