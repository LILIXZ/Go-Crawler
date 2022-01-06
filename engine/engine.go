package engine

import (
	"GO-CRAWLER/fetcher"
	"fmt"
)

type Engine struct {
	Scheduler Scheduler
}

type Scheduler interface {
	Submit(Request)
	ConfigureScheduler(chan Request)
}

const initializeWorkerCount = 10

func (e Engine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureScheduler(in)

	for i := 0; i < initializeWorkerCount; i++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		if !isDuplicated(r.Url) {
			e.Scheduler.Submit(r)
		}
		// else {
		// 	fmt.Println("Duplicated URL:", r.Url)
		// }
	}

	for {
		result := <-out

		for _, item := range result.Items {
			if article, ok := item.(Article); ok {
				fmt.Printf("Got article: %s", article.Heading)
				fmt.Println()
			} else {
				fmt.Printf("Got item: %s", item)
				fmt.Println()
			}
		}
		for _, request := range result.Requests {
			if !isDuplicated(request.Url) {
				e.Scheduler.Submit(request)
			}
			// else {
			// 	fmt.Println("Duplicated URL:", request.Url)
			// }
		}
	}

}

func worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)

	fmt.Println("URL:", r.Url)

	if err != nil {
		fmt.Printf("Error occurred: %s", err)

		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedURL = make(map[string]bool)

func isDuplicated(url string) bool {
	if visitedURL[url] {
		return true
	}
	visitedURL[url] = true
	return false
}
