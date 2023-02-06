package web

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

func Request(uri string) ([]byte, error) {
	req := fasthttp.AcquireRequest()

	req.SetRequestURI(uri)
	req.Header.SetCookie("CONSENT", "YES+")
	req.Header.SetUserAgent(uuid.NewString())
	resp := fasthttp.AcquireResponse()
    defer fasthttp.ReleaseResponse(resp)

	err := fasthttp.Do(req, resp)

	if err != nil {
		return nil, err
	}

	contentEncoding := resp.Header.Peek("Content-Encoding")
    var body []byte
    if bytes.EqualFold(contentEncoding, []byte("gzip")) {
        body, _ = resp.BodyGunzip()
    } else {
        body = resp.Body()
    }

	return body, nil
}

func RequestJson(uri string, v any) error {
	body, err := Request(uri)

	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &v)

	return err
}

func RequestHtml(uri string) (*goquery.Document, error) {
	body, err := Request(uri)

	if err != nil {
		return nil, err
	}

	reader := strings.NewReader(string(body))
    doc, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		return nil, err
	}
	return doc, nil
}