package enums
//This package contains enums which will be commonly used by all packages within website-scraping

//streaming services names will be stored by this data type
type StreamingServiceName string

const (
	Amazon = StreamingServiceName("amazon")
)

//parser types will be stored by this data type. It will be used to determine which parser to be used by a streaming service
type ParserType string

const (
	AmazonParser1    = ParserType("amazon-parser-1")
	AmazonMockParser = ParserType("amazon-mock-parser")
)

