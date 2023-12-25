package http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func Get[T struct{}](rawUrl string, params map[string]string) (*T, error) {
	fUrl, err := constructUrl(rawUrl, params)
	if err != nil {
		return nil, err
	}

	fmt.Println("GET", fUrl)
	res, err := http.Get(fUrl)
	if err != nil {
		return nil, err
	}

	body, err := parseBody[T](res)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func Post[T struct{}, S struct{}](urlPath string, request *S) (*T, error) {
	fmt.Println("POST", urlPath)

	requestJson, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(requestJson)
	res, err := http.Post(urlPath, "application/json", reader)
	if err != nil {
		return nil, err
	}

	body, err := parseBody[T](res)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func parseBody[T struct{}](res *http.Response) (*T, error) {
	body := &T{}
	return body, json.NewDecoder(res.Body).Decode(body)
}

func constructUrl(rawUrl string, params map[string]string) (string, error) {
	fUrl, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	urlValues := &url.Values{}
	for k, v := range params {
		urlValues.Add(k, v)
	}
	fUrl.RawQuery = urlValues.Encode()
	return fUrl.String(), nil
}
