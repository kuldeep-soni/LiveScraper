package httpClient

import (
	"context"
	"io"
	"net/http"
)

type httpClient struct {
}

func (h *httpClient) GetHTMLResource(ctx context.Context, url string) (response string, err error) {
	headers := map[string]string{
		"accept": "text/html",
	}

	req, _ := http.NewRequest("GET", url, nil)
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	dataBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return string(dataBytes), nil
}

func newHttpClient() *httpClient {
	return &httpClient{}
}