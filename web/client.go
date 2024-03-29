package web

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Request(uri string) ([]byte, []byte, error) {
	client := *http.DefaultClient
	req, err := http.NewRequest("GET", uri, nil)

	if err != nil {
		return nil, nil, err
	}

	req.AddCookie(&http.Cookie{Name: "CONSENT", Value: "YES+"})
	
	if (!strings.Contains(uri, "www.google.com")) {
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	}

	resp, err := client.Do(req)

	if err != nil || resp.StatusCode == 404 {
		return nil, nil, err
	}

	contentType := resp.Header.Get("Content-Type")

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, nil, err
	}

	contentEncoding := resp.Header.Get("Content-Encoding")
	if bytes.EqualFold([]byte(contentEncoding), []byte("gzip")) {
		body, _ = gUnzipData(body)
	}

	return body, []byte(contentType), nil
}

func RequestJson(uri string, v any) error {
	body, _, err := Request(uri)

	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &v)

	return err
}

func RequestHtml(uri string) (*goquery.Document, error) {
	body, _, err := Request(uri)

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

func gUnzipData(data []byte) (resData []byte, err error) {
	b := bytes.NewBuffer(data)

	var r io.Reader
	r, err = gzip.NewReader(b)
	if err != nil {
		return
	}

	var resB bytes.Buffer
	_, err = resB.ReadFrom(r)
	if err != nil {
		return nil, err
	}

	resData = resB.Bytes()

	return
}
