package phttp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DataResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Success bool        `json:"success"`
}

type ErrorResponse struct {
	Success bool      `json:"success"`
	Error   ErrorBody `json:"error"`
}

type ErrorBody struct {
	Code     string                 `json:"code"`
	Message  string                 `json:"message"`
	MetaData map[string]interface{} `json:"metadata,omitempty"`
}

func SendPrettifiedSuccessJsonResponse(c *gin.Context, response interface{}) {
	c.IndentedJSON(http.StatusOK, response)
}

func SendFailureResponse(c *gin.Context, httpCode int, code, message string) {
	c.JSON(httpCode, ErrorResponse{
		Success: false,
		Error: ErrorBody{
			Code:    code,
			Message: message,
		},
	})
}
