package config

import (
	"encoding/json"
	"flag"

	_ "embed"

	"github.com/bnyro/findx/entities"
)

var Addr *string
var Proxy *bool
var Redirects []entities.Redirect

//go:embed redirects.json
var redirectsFile []byte

func Init() {
	Addr = flag.String("addr", ":8080", "address to listen on")
	Proxy = flag.Bool("proxy", false, "weather to proxy all images")

	json.Unmarshal(redirectsFile, &Redirects)

	flag.Parse()
}
