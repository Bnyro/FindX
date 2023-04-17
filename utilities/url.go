package utilities

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/bnyro/findx/config"
)

var numeric = regexp.MustCompile(`\d`)
var urlChars = regexp.MustCompile(`[_-]`)

func RewriteProxied(uri string) string {
	if !*config.Proxy {
		return uri
	}
	target := url.QueryEscape(uri)
	return fmt.Sprintf("/proxy?url=%s", target)
}

func Redirect(uri string) string {
	parsedUrl, err := url.Parse(uri)

	if err != nil {
		return uri
	}

	for _, redirect := range config.Redirects {
		if parsedUrl.Hostname() == redirect.Source {
			return strings.Replace(uri, redirect.Source, redirect.Target, 1)
		}
	}

	return uri
}

func HumanizeUrl(uri string) string {
	parsedUrl, err := url.Parse(uri)
	if err != nil {
		return uri
	}
	readableUrl := parsedUrl.Host
	pathComponents := strings.Split(parsedUrl.Path, "/")
	for index, path := range pathComponents {
		if len(path) == 2 || IsBlank(path) || numeric.MatchString(path) {
			// most likely a language code, e.g. de, or just blank in general
			continue
		}
		if len(pathComponents)-1 == index && strings.Contains(path, ".") {
			// skip the last path component if it's a file
			continue
		}
		component := urlChars.ReplaceAllString(path, " ")
		readableUrl = fmt.Sprintf("%s › %s", readableUrl, strings.Title(component))
	}
	return readableUrl
}
