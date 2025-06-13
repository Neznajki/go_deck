package main

import (
	"fmt"
	"net/http"
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
	}

	channel := make(chan string, len(links))

	for _, link := range links {
		go checkSite(link, channel)
	}

	for range links {
		fmt.Println(<-channel)
	}
}

func checkSite(link string, channel chan string) {
	_, err := http.Get(link)

	if err != nil {
		channel <- fmt.Sprintf("Link Not Working : %s, with reason %v", link, err.Error())
		return
	}

	channel <- fmt.Sprintf("Link Working : %s", link)
}
