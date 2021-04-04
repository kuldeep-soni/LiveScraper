package model

import (
	"errors"
	"fmt"
)

type MovieMeta struct {
	Title       string   `json:"title"`
	ReleaseYear int      `json:"release_year"`
	Actors      []string `json:"actors"`
	Poster      string   `json:"poster"`
	SimilarIds  []string `json:"similar_ids"`
}

type Source string

const (
	Amazon = Source("amazon")
)

//move url to config
func (s Source) GetWebsiteUrl() (url string, err error) {
	switch s {
	case Amazon:
		url = "http://www.amazon.de/gp/product"
	default:
		err = errors.New(fmt.Sprintf("GetWebsiteUrl: Given source url has not been set, source: %v", s))
	}
	return
}
