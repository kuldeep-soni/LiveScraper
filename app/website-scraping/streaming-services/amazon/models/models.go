package models

import "github.com/LiveScraper/app/website-scraping/model"

type MovieMetaAmazon struct {
	Title       string   `json:"title"`
	ReleaseYear int      `json:"release_year"`
	Actors      []string `json:"actors"`
	Poster      string   `json:"poster"`
	SimilarIds  []string `json:"similar_ids"`
}

func (mma MovieMetaAmazon) GetGlobalMovieMeta() (mm model.MovieMeta) {
	mm.Title = mma.Title
	mm.ReleaseYear = mma.ReleaseYear
	mm.Actors = mma.Actors
	mm.Poster = mma.Poster
	mm.SimilarIds = mma.SimilarIds
	return
}
