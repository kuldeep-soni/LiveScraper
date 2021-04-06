package model
//All the structs which will be used commonly by all the packages within website-scraping app
//are defined here

//MovieMeta is a structure which contains the desired output to be fetched from all streaming services
type MovieMeta struct {
	Title       string   `json:"title"`
	ReleaseYear int      `json:"release_year"`
	Actors      []string `json:"actors"`
	Poster      string   `json:"poster"`
	SimilarIds  []string `json:"similar_ids"`
}