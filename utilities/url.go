package utilities

import (
	"fmt"
	"net/url"

	"github.com/bnyro/findx/config"
)

func RewriteProxied(uri string) string {
	if !*config.Proxy {
		return uri
	}
	target := url.QueryEscape(uri)
	return fmt.Sprintf("/proxy?url=%s", target)
}
