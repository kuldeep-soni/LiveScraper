package httpClient

import "context"

type IHttpClient interface {
	GetHTMLResource(ctx context.Context, url string) (response string, err error)
}
