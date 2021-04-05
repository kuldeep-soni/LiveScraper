package global

var GStreamingServices *StreamingServices

type StreamingServices []StreamingService

type StreamingService struct {
	Name string `json:"name"`
	URL string `json:"url"`
	ParserType string `json:"parser_type"`
}
