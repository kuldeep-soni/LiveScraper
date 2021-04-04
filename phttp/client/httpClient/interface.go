package httpClient

import (
	"context"
)

type IHttpClient interface {
	GetHTMLResource(ctx context.Context, url string) (response string, err error)
}

func GetIHttpClient(implType string, userAgents []string) IHttpClient {
	switch implType {
	case "actual":
		return newHttpClient(userAgents)
	default:
		return newMockHttpClient()
	}
}
