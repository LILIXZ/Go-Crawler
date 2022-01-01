package parser

import (
	"GO-CRAWLER/engine"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var articleRe = regexp.MustCompile(`(?s)<article(.+)</article`)                                       // extract match
var headingRe = regexp.MustCompile(`(?U)id="main-heading"[^>]*>([^<]*)</`)                            // extract match[0][1]
var summaryRe = regexp.MustCompile(`(?mU)data-component="text-block".*BoldText[^>]*>([^<]*)</b></p>`) // extract match[0][1]
var contentRe = regexp.MustCompile(`(?mU)data-component="(text-block|image-block|unordered-list-block)".*(Paragraph[^>]*>(.*)</p>|BoldText[^>]*>([^<]*)</b>|srcSet="(\S*?)|<ul[^>]*>(.*)</ul>)`)

func ParseArticle(contents []byte) engine.ParseResult {
	// regex match article
	matchArticle := articleRe.FindAllStringSubmatch(string(contents), -1)
	article := ""

	if len(matchArticle) > 0 && len(matchArticle[0]) >= 2 {
		article = matchArticle[0][1]
	}

	// regex match title
	matchTitle := headingRe.FindAllStringSubmatch(article, -1)

	if len(matchTitle) > 0 && len(matchTitle[0]) >= 2 {
		fmt.Println("Title:", matchTitle[0][1])
	}

	// regex match summary
	matchSummary := summaryRe.FindAllStringSubmatch(article, -1)

	if len(matchSummary) > 0 && len(matchSummary[0]) >= 2 {
		fmt.Println("Summary:", matchSummary[0][1])
	}

	// regex match contents
	matchContents := contentRe.FindAllStringSubmatch(article, -1)
	fmt.Println("---------- Contents ----------")
	for _, m := range matchContents {
		if len(m) >= 2 {
			fmt.Println(m[1])

		}
	}

	if strings.Contains(matchTitle[0][1], "India") {
		f, _ := os.Create("bbc_content3.txt")
		for _, m := range matchContents {
			fmt.Fprintln(f, m[1]) // block type
			fmt.Fprintln(f)
			if len(m) >= 3 {
				fmt.Fprintln(f, m[3]) // text
			}
			if len(m) >= 4 {
				fmt.Fprintln(f, m[4])
			}
			if len(m) >= 5 {
				fmt.Fprintln(f, m[5]) // image link
			}
			if len(m) >= 6 {
				fmt.Fprintln(f, m[6]) // unordered list
			}
		}

		f.Close()
	}

	return engine.ParseResult{}
}
