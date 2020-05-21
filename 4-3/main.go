package main

import "fmt"

func main() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			for s := range strings {
				fmt.Println(s)
			}
		}()
		return completed
	}

	doWork(nil)
	fmt.Println("Done.")
}
