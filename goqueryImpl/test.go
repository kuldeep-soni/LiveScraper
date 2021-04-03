package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Response struct {
	Title string `json:"title"`
	ReleaseYear string `json:"release_year"`
	Actors []string `json:"actors"`
	Poster string `json:"poster"`
	SimilarIds []string `json:"similar_ids"`
}

func ScrapeTextFile(doc *goquery.Document, response *Response){
	seletion := doc.Find(".DVWebNode-detail-atf-wrapper.DVWebNode")

	title := seletion.Find("._2IIDsE._3I-nQy")
	dataAutomationId, exists := title.Attr("data-automation-id")
	if exists && dataAutomationId == "title"{
		response.Title = title.Text()
	}

	metaInfo := seletion.Find("._2IIDsE._30UT8H")
	dataAutomationId, exists = metaInfo.Attr("data-automation-id")
	if exists && dataAutomationId == "meta-info"{
		_ = metaInfo.Find("dl").Each(func(i int, s *goquery.Selection) {
			if s.Find("dt").Text() == "Hauptdarsteller"{
				s.Find("dd").Find("a").Each(func(j int, p *goquery.Selection) {
					response.Actors = append(response.Actors, p.Text())
				})
			}
		})
	}

	releaseYear := seletion.Find(".dv-node-dp-badges").ChildrenFiltered("span").FilterFunction(func(i int, s *goquery.Selection) (isSelected bool){
		s.Find("span").Each(func (i int, s *goquery.Selection){
			releaseYear, exists := s.Attr("data-automation-id")
			if exists && releaseYear == "release-year-badge"{
				isSelected = true
			}
		})
		return
	})
	response.ReleaseYear = releaseYear.Text()

	relatedMoviesList := doc.Find(".DVWebNode-detail-btf-wrapper.DVWebNode").Find("div[data-automation-id=caw-carousel-section]").Find("div.jxBPRE.mWsquZ").Find("ul").Children()
	relatedMoviesList.Each(func(i int, s *goquery.Selection){
		movieLink, ok := s.Find("div._1Opa2_").Find("a").Attr("href")
		if ok{
			temp := strings.Split(movieLink, "/")
			if len(temp) >= 5{
				response.SimilarIds = append(response.SimilarIds, temp[4])
			}
		}
	})

	posterComponent := seletion.Find("#atf-full")
	posterLink, ok := posterComponent.Attr("src")
	if ok{
		response.Poster = posterLink
	}
}

func main() {
	response := Response{}
	//B00KY1U7GM, B00K19SD8Q, B08MDJPYD9, B08RYBTG7S, B08FMQTK65, B00FCM7N9C
	dataBytes, err := ioutil.ReadFile("data/B00FCM7N9C.txt")
	if err != nil{
		fmt.Print(err)
	}

	readerPointer := strings.NewReader(string(dataBytes))
	doc, err := goquery.NewDocumentFromReader(readerPointer)
	if err != nil {
		log.Fatal(err)
	}

	ScrapeTextFile(doc, &response)
	responseBytes, _ := json.Marshal(response)
	var jsonResponse map[string]interface{}
	json.Unmarshal(responseBytes, &jsonResponse)
	fmt.Println(jsonResponse)
}