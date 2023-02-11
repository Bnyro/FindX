package config

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/bnyro/findx/entities"
)

var Addr *string
var Proxy *bool
var Redirects []entities.Redirect

func Init() {
	Addr = flag.String("addr", ":8080", "address to listen on")
	Proxy = flag.Bool("proxy", false, "weather to proxy all images")

	redirectsFile, err := os.ReadFile("./redirects.json")
	if err == nil {
		json.Unmarshal(redirectsFile, &Redirects)
	}

	flag.Parse()
}
