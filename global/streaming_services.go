package global

var GStreamingServices *StreamingServices

type StreamingServices []StreamingService

//StreamingService struct represents a unique api endpoint that returns an html document
//It is expected to have same DOM structure for every resource id passed to it
//Example: Amazon Prime and Netflix will be different streaming services because their html doc has
//different structure
//
//Name is defined in configuration.json. Needed as an identifier
//URL is a unique api endpoint corresponding to the StreamingService
//ParserType defines what internal implementation of DOM parsing are we going to use
type StreamingService struct {
	Name string `json:"name"`
	URL string `json:"url"`
	ParserType string `json:"parser_type"`
}
