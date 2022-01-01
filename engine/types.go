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
	Heading  string
	Contents []Content
}

type Content struct {
	Paragraph, Summary, ImageUrl string
	IsImage                      bool
	Lists                        []string
}

func NilParser(contents []byte) ParseResult {
	return ParseResult{}
}
