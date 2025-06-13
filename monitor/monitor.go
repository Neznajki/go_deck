package main

import (
	"fmt"
	"net/http"
	"slices"
	"time"
)

type RequestData struct {
	Link      string
	c         chan RequestData
	terminate bool
	tries     int
}

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

func executeWithRetry(rd RequestData) {
	if rd.tries >= 10 {
		rd.terminate = true
		fmt.Println(fmt.Sprintf("done requests due to limitation for link (%s)", rd.Link))
		rd.c <- rd
		return
	}

	rd = monitorSite(rd)
	if rd.terminate {
		rd.c <- rd
		return
	}
	pause(2 * time.Second)

	fmt.Println("--retry link-- ", rd.Link)
	rd.c <- rd
}

func monitorSite(rd RequestData) RequestData {
	_, err := http.Get(rd.Link)

	rd.tries++

	if err != nil {
		fmt.Println(fmt.Sprintf("link Not Working : %s, with reason %v", rd.Link, err.Error()))
		rd.terminate = true
		return rd
	}

	fmt.Println("link Working : ", rd.Link)
	return rd
}

func pause(timeToPause time.Duration) {
	time.Sleep(timeToPause)
}
