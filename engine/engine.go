package engine

import (
	"GO-CRAWLER/fetcher"
	"fmt"
)

func Run(seeds ...Request) {
	var requests []Request

	// Append seeds to the request queue
	requests = append(requests, seeds...)

	// Looping request queue, fetch the content and give to the parser
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		body, err := fetcher.Fetch(r.Url)

		fmt.Println("URL:", r.Url)

		if err != nil {
			fmt.Printf("Error occurred: %s", err)
			continue
		}

		parseRequest := r.ParserFunc(body)
		requests = append(requests, parseRequest.Requests...)

		for _, item := range parseRequest.Items {
			fmt.Printf("Got item: %s", item)
			fmt.Println()
		}
	}

}
