package httpClient

import (
	"context"
	"io"
	"math/rand"
	"net/http"
)

type httpClient struct {
	userAgents []string
}

func (h *httpClient) GetHTMLResource(ctx context.Context, url string) (response string, err error) {
	userAgent := h.userAgents[rand.Intn(len(h.userAgents))]
	headers := map[string]string{
		"accept":     "text/html",
		"user-agent": userAgent,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
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

func newHttpClient(userAgents []string) *httpClient {
	return &httpClient{
		userAgents: userAgents,
	}
}