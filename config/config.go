package config

import "flag"

var Addr *string
var Proxy *bool

func Init() {
	Addr = flag.String("addr", ":8080", "address to listen on")
	Proxy = flag.Bool("proxy", false, "weather to proxy all images")
	flag.Parse()
}
