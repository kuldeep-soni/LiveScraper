package enums

type StreamingServiceName string

const (
	Amazon = StreamingServiceName("amazon")
)

type ParserType string

const (
	AmazonParser1    = ParserType("amazon-parser-1")
	AmazonMockParser = ParserType("amazon-mock-parser")
)

