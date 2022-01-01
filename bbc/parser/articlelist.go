package parser

import (
	"GO-CRAWLER/engine"
	"regexp"
)

var articleListRe = regexp.MustCompile(`href="(/news/[^"]+)">[^<]*<h3[^>]*>([^<]+)<`)

const domainUrl = "https://www.bbc.com"

func ParseArticleList(contents []byte) engine.ParseResult {

	// re := regexp.MustCompile(`href="(/news/[^"]+)">.*title[^>]*>([^<]+)<`)
	matches := articleListRe.FindAllStringSubmatch(string(contents), -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: domainUrl + m[1], ParserFunc: ParseArticle,
		})
	}
	return result
}
