package scraper

import (
	"context"
	"fmt"
	"github.com/LiveScraper/app/website-scraping/scraper/model"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type IService interface {
	GetMovieMeta(ctx context.Context, movieId string) (model.MovieMeta, error)
}

type service struct {
}

func (s *service) GetMovieMeta(ctx context.Context, movieId string) (meta model.MovieMeta, err error) {
	dataBytes, err := ioutil.ReadFile("data/B00FCM7N9C.txt")
	if err != nil {
		fmt.Print(err)
	}

	readerPointer := strings.NewReader(string(dataBytes))
	doc, err := goquery.NewDocumentFromReader(readerPointer)
	if err != nil {
		log.Fatal(err)
	}

	scrapeTextFile(doc, &meta)
	return
}

func scrapeTextFile(doc *goquery.Document, response *model.MovieMeta) {
	seletion := doc.Find(".DVWebNode-detail-atf-wrapper.DVWebNode")

	title := seletion.Find("._2IIDsE._3I-nQy")
	dataAutomationId, exists := title.Attr("data-automation-id")
	if exists && dataAutomationId == "title" {
		response.Title = title.Text()
	}

	metaInfo := seletion.Find("._2IIDsE._30UT8H")
	dataAutomationId, exists = metaInfo.Attr("data-automation-id")
	if exists && dataAutomationId == "meta-info" {
		_ = metaInfo.Find("dl").Each(func(i int, s *goquery.Selection) {
			if s.Find("dt").Text() == "Hauptdarsteller" {
				s.Find("dd").Find("a").Each(func(j int, p *goquery.Selection) {
					response.Actors = append(response.Actors, p.Text())
				})
			}
		})
	}

	releaseYear := seletion.Find(".dv-node-dp-badges").ChildrenFiltered("span").FilterFunction(func(i int, s *goquery.Selection) (isSelected bool) {
		s.Find("span").Each(func(i int, s *goquery.Selection) {
			releaseYear, exists := s.Attr("data-automation-id")
			if exists && releaseYear == "release-year-badge" {
				isSelected = true
			}
		})
		return
	})
	response.ReleaseYear, _ = strconv.Atoi(releaseYear.Text())

	relatedMoviesList := doc.Find(".DVWebNode-detail-btf-wrapper.DVWebNode").Find("div[data-automation-id=caw-carousel-section]").Find("div.jxBPRE.mWsquZ").Find("ul").Children()
	relatedMoviesList.Each(func(i int, s *goquery.Selection) {
		movieLink, ok := s.Find("div._1Opa2_").Find("a").Attr("href")
		if ok {
			temp := strings.Split(movieLink, "/")
			if len(temp) >= 5 {
				response.SimilarIds = append(response.SimilarIds, temp[4])
			}
		}
	})

	posterComponent := seletion.Find("#atf-full")
	posterLink, ok := posterComponent.Attr("src")
	if ok {
		response.Poster = posterLink
	}
}

func NewService() IService{
	return &service{}
}