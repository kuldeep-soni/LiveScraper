package amazon

import (
	"context"
	"github.com/LiveScraper/app/website-scraping/scraper/model"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

type amazonDocumentParser1 struct {
}

func GetAmazonDocumentParser1() *amazonDocumentParser1 {
	return &amazonDocumentParser1{}
}

func (a *amazonDocumentParser1) TransformRawMovieData(ctx context.Context, rawData string) (meta model.MovieMeta, err error) {
	////B00KY1U7GM, B00K19SD8Q, B08MDJPYD9, B08RYBTG7S, B08FMQTK65, B00FCM7N9C
	readerPointer := strings.NewReader(rawData)
	doc, err := goquery.NewDocumentFromReader(readerPointer)
	if err != nil {
		return
	}

	return a.transformRawMovieData(ctx, doc)
}

func (a *amazonDocumentParser1) transformRawMovieData(ctx context.Context, doc *goquery.Document) (meta model.MovieMeta, err error) {
	selection := doc.Find(".DVWebNode-detail-atf-wrapper.DVWebNode")

	title := selection.Find("._2IIDsE._3I-nQy")
	dataAutomationId, exists := title.Attr("data-automation-id")
	if exists && dataAutomationId == "title" {
		meta.Title = title.Text()
	}

	metaInfo := selection.Find("._2IIDsE._30UT8H")
	dataAutomationId, exists = metaInfo.Attr("data-automation-id")
	if exists && dataAutomationId == "meta-info" {
		_ = metaInfo.Find("dl").Each(func(i int, s *goquery.Selection) {
			if s.Find("dt").Text() == "Hauptdarsteller" {
				s.Find("dd").Find("a").Each(func(j int, p *goquery.Selection) {
					meta.Actors = append(meta.Actors, p.Text())
				})
			}
		})
	}

	releaseYear := selection.Find(".dv-node-dp-badges").ChildrenFiltered("span").FilterFunction(func(i int, s *goquery.Selection) (isSelected bool) {
		s.Find("span").Each(func(i int, s *goquery.Selection) {
			releaseYear, exists := s.Attr("data-automation-id")
			if exists && releaseYear == "release-year-badge" {
				isSelected = true
			}
		})
		return
	})
	meta.ReleaseYear, _ = strconv.Atoi(releaseYear.Text())

	relatedMoviesList := doc.Find(".DVWebNode-detail-btf-wrapper.DVWebNode").Find("div[data-automation-id=caw-carousel-section]").Find("div.jxBPRE.mWsquZ").Find("ul").Children()
	relatedMoviesList.Each(func(i int, s *goquery.Selection) {
		movieLink, ok := s.Find("div._1Opa2_").Find("a").Attr("href")
		if ok {
			temp := strings.Split(movieLink, "/")
			if len(temp) >= 5 {
				meta.SimilarIds = append(meta.SimilarIds, temp[4])
			}
		}
	})

	posterComponent := selection.Find("#atf-full")
	posterLink, ok := posterComponent.Attr("src")
	if ok {
		meta.Poster = posterLink
	}

	return
}
