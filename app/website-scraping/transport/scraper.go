package transport

import (
	"github.com/LiveScraper/app/website-scraping/enums"
	"github.com/LiveScraper/app/website-scraping/service"
	"github.com/LiveScraper/phttp"
	"github.com/gin-gonic/gin"
)

//This is the transport layer of scraper service
//StreamingServiceHandler listens to http requests made by clients at endpoints listed here
func StreamingServiceHandler(r *gin.RouterGroup, s service.IScraper) {
	r.GET("amazon/:amazonId", getMovieMeta(s, enums.Amazon))
}

//It fetches parsed data in a well defined structure and sends it as a JSON response
//This can be used by any StreamingService
func getMovieMeta(scraperSvc service.IScraper, ssName enums.StreamingServiceName) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		movieId := ctx.Param("amazonId")

		meta, err := scraperSvc.GetMovieMeta(ctx, ssName, movieId)
		if err != nil {
			phttp.SendFailureResponse(ctx, 500, "ISE500", "Something went wrong")
			return
		}

		phttp.SendPrettifiedSuccessJsonResponse(ctx, meta)
	}
}