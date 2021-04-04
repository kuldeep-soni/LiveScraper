package httpClient

import (
	"context"
)

type IHttpClient interface {
	GetHTMLResource(ctx context.Context, url string) (response string, err error)
}

func GetIHttpClient(implType string) IHttpClient {
	switch implType {
	case "actual":
		return newHttpClient()
	default:
		return newMockHttpClient()
	}
}
