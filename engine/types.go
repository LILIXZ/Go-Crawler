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
	Title    string
	Contents []Content
}

type Content struct {
	Paragraph, Summary, ImageUrl string
	IsImage                      bool
}

func NilParser(contents []byte) ParseResult {
	return ParseResult{}
}
