package main

import (
	"fmt"
	"net/http"
	"slices"
	"time"
)

type RequestData struct {
	Link   string
	c      chan RequestData
	repeat bool
	tries  int
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
	}

	channel := make(chan RequestData, len(links))

	for _, link := range links {
		rd := RequestData{Link: link, c: channel, repeat: true}
		go monitorSite(rd)
	}

	for len(links) > 0 {
		rd := <-channel
		if rd.repeat {
			go monitorSite(rd)
			continue
		}

		links = slices.DeleteFunc(links, func(s string) bool {
			return s == rd.Link
		})

		fmt.Println("----done monitoring---- ", rd.Link)
	}
}

func monitorSite(rd RequestData) {
	_, err := http.Get(rd.Link)

	rd.tries++

	if rd.tries >= 10 {
		rd.repeat = false
		fmt.Println(fmt.Sprintf("done requests due to limitation for link (%s)", rd.Link))
	}

	if err != nil {
		fmt.Println(fmt.Sprintf("link Not Working : %s, with reason %v", rd.Link, err.Error()))
		rd.repeat = false
		rd.c <- rd
		return
	}

	fmt.Println("link Working : ", rd.Link)
	pause(2 * time.Second)

	fmt.Println("--retry link-- ", rd.Link)
	rd.c <- rd
}

func pause(timeToPause time.Duration) {
	time.Sleep(timeToPause)
}
