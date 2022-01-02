package parser

import (
	"GO-CRAWLER/engine"
	"html"
	"regexp"
)

var articleRe = regexp.MustCompile(`(?s)<article(.+)</article`)                                       // extract match
var headingRe = regexp.MustCompile(`(?U)id="main-heading"[^>]*>([^<]*)</`)                            // extract match[0][1]
var summaryRe = regexp.MustCompile(`(?mU)data-component="text-block".*BoldText[^>]*>([^<]*)</b></p>`) // extract match[0][1]
var contentRe = regexp.MustCompile(`(?mU)data-component="(text-block|image-block|unordered-list-block|subheadline-block)".*(Paragraph[^>]*>(.*)</p>|BoldText[^>]*>(.*)</b>|srcSet="(\S*?)|<ul[^>]*>(.*)</ul>|<span role="text">(.*)</span>)`)
var seeAlsoRe = regexp.MustCompile(`data-component="see-alsos"(.*)</section>`)
var urlRe = regexp.MustCompile(`href="(/news/[^"]+)"`)

func ParseArticle(contents []byte) engine.ParseResult {
	result := engine.ParseResult{
		Items:    nil,
		Requests: []engine.Request{},
	}
	articleItem := engine.Article{Heading: "", Summary: "", Contents: []engine.Content{}}
	// regex match article
	matchArticle := articleRe.FindAllStringSubmatch(string(contents), -1)
	article := ""

	if len(matchArticle) > 0 && len(matchArticle[0]) >= 2 {
		article = matchArticle[0][1]
	}

	// regex - title
	matchTitle := headingRe.FindAllStringSubmatch(article, -1)

	if len(matchTitle) > 0 && len(matchTitle[0]) >= 2 {
		articleItem.Heading = html.UnescapeString(matchTitle[0][1])
	}

	// regex - summary
	matchSummary := summaryRe.FindAllStringSubmatch(article, -1)

	if len(matchSummary) > 0 && len(matchSummary[0]) >= 2 {
		articleItem.Summary = html.UnescapeString(matchSummary[0][1])
	}

	// regex - contents
	matchContents := contentRe.FindAllStringSubmatch(article, -1)
	for _, m := range matchContents {
		contentItem := engine.Content{}
		if len(m) >= 2 {
			contentItem.Type = m[1] // block type
		}
		if len(m) > 3 && len(m[3]) > 0 {
			contentItem.Paragraph = html.UnescapeString(m[3]) // text
		}
		if len(m) > 5 && len(m[5]) > 0 {
			contentItem.ImageUrl = html.UnescapeString(m[5]) // image link
		}
		if len(m) > 6 && len(m[6]) > 0 {
			contentItem.Paragraph = html.UnescapeString(m[6]) // unordered list
		}
		if len(m) > 7 && len(m[7]) > 0 {
			contentItem.Paragraph = html.UnescapeString(m[7]) // subheadline content
		}

		articleItem.Contents = append(articleItem.Contents, contentItem)
	}

	result.Items = append(result.Items, articleItem)

	// regex - see alsos
	matchSeeAlsos := seeAlsoRe.FindAllStringSubmatch(article, -1)
	seeAlsos := ""
	if len(matchSeeAlsos) > 0 && len(matchSeeAlsos[0]) >= 2 {
		seeAlsos = matchSeeAlsos[0][1]
	}

	// regex - see alsos URLs
	matchURLs := urlRe.FindAllStringSubmatch(seeAlsos, -1)
	if len(matchURLs) > 0 {
		for _, m := range matchURLs {
			if len(m) >= 2 {
				result.Requests = append(result.Requests, engine.Request{Url: domainUrl + m[1], ParserFunc: ParseArticle})
			}
		}
	}

	return result
}
