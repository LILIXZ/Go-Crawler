package main

import (
	"GO-CRAWLER/bbc/parser"
	"GO-CRAWLER/engine"
)

func main() {
	engine.Run(engine.Request{Url: "https://www.bbc.com/news", ParserFunc: parser.ParseArticleList})
}
