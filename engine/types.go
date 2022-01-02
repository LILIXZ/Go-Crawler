package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

type Article struct {
	Heading, Summary string
	Contents         []Content
}

type Content struct {
	Type, Paragraph, ImageUrl string
}

func NilParser(contents []byte) ParseResult {
	return ParseResult{}
}
