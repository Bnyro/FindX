package utilities

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/bnyro/findx/config"
)

func RewriteProxied(uri string) string {
	if !*config.Proxy {
		return uri
	}
	target := url.QueryEscape(uri)
	return fmt.Sprintf("/proxy?url=%s", target)
}

func Redirect(uri string) string {
	parsedUrl, _ := url.Parse(uri)

	for _, redirect := range config.Redirects {
		if parsedUrl.Hostname() == redirect.Source {
			return strings.Replace(uri, redirect.Source, redirect.Target, 1)
		}
	}

	return uri
}
