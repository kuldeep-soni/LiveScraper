package scraper

import (
	"fmt"
	"github.com/LiveScraper/phttp"
	"github.com/gin-gonic/gin"
)

//how to handle concurrency
//we want small, synchronous, iterator based, fault tolerant with zero external dependencies based html parser
//Amazon can temporarily block the IP from which automated requests go. Different means can be used for it.
//For example, Amazon may show a captcha or a page with an error. Therefore, for the scraper to work successfully,
//we need to think about how it will catch and bypass these cases.
//can we use go channels
//it is not guaranteed that the io.Writer is safe to use concurrent
//try accessing data using proxies
//what is user agent in headers
//running benchmarks
//can we use regex for this?
//css selector vs XPath
//remove all /n, /t
//make a complete path of the document
//what other libraries can be used?
//extendable
//web server
//scraping service -> scrape based on which website to scrape, which location to scrape, have a config file which can define
//all the css selectors so that redeploy can just suffice
//what if new data is to be included?
//including a service which escapes

func Handler(r *gin.RouterGroup, s IService) {
	r.GET("movie/amazon/:amazonId", getAmazonMovieMeta(s))
}

//how to limit rate of this api?
func getAmazonMovieMeta(s IService) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		amazonId := ctx.Param("amazonId")
		fmt.Println(amazonId)

		meta, err := s.GetMovieMeta(ctx, amazonId)
		if err != nil{
			phttp.SendFailureResponse(ctx, 500, "ISE500", "Something went wrong")
			return
		}

		phttp.SendPrettifiedSuccessJsonResponse(ctx, meta)
	}
}