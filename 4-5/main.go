package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Error    error
	Response *http.Response
}

func main() {
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		resulsts := make(chan Result)
		go func() {
			defer close(resulsts)

			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp}
				select {
				case <-done:
					return
				case resulsts <- result:
				}
			}
		}()
		return resulsts
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://google.com", "https://alitaso345"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error: %v\n", result.Error)
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}
