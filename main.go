package main

import (
	"GO-CRAWLER/bbc/parser"
	"GO-CRAWLER/engine"
	"GO-CRAWLER/scheduler"
)

func main() {
	e := engine.Engine{Scheduler: &scheduler.Scheduler{}}
	e.Run(engine.Request{Url: "https://www.bbc.com/news", ParserFunc: parser.ParseArticleList})
}
