package utilities

import (
	"fmt"
	"net/url"
)

func RewriteProxied(uri string) string {
	target := url.QueryEscape(uri)
	return fmt.Sprintf("/proxy?url=%s", target)
}
