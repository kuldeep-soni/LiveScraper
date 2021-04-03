package scraper

import (
	"fmt"
	"github.com/LiveScraper/phttp"
	"github.com/gin-gonic/gin"
)

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