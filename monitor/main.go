package main

import (
	"fmt"
	"slices"
)

func main() {
	links := []string{
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.google.com",
		"https://www.bing.com",
		"https://www.bing.com.cn",
		"https://www.amazon.com",
		"https://www.microsoft.cn",
		"https://www.microsoft.com",
	}

	channel := make(chan RequestData, len(links))

	for _, link := range links {
		rd := RequestData{Link: link, c: channel}
		go executeWithRetry(rd)
	}

	for len(links) > 0 {
		rd := <-channel
		if rd.terminate {
			links = slices.DeleteFunc(links, func(s string) bool {
				return s == rd.Link
			})

			fmt.Println("----done monitoring---- ", rd.Link)
			continue
		}

		go executeWithRetry(rd)
	}
}
