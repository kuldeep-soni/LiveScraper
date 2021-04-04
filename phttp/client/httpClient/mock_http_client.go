package httpClient

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"
)

type mockHttpClient struct{}

func (m *mockHttpClient) GetHTMLResource(ctx context.Context, url string) (response string, err error) {
	splitUrl := strings.Split(url, "/")
	productId := splitUrl[len(splitUrl)-1]

	responseBytes, err := ioutil.ReadFile(fmt.Sprintf("data/%v.txt", productId))
	if err != nil {
		return
	}

	return string(responseBytes), nil
}

func newMockHttpClient() *mockHttpClient {
	return &mockHttpClient{}
}
