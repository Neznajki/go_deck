package main

import (
	"fmt"
	"net/http"
)

func checkSite(link string, channel chan string) {
	_, err := http.Get(link)

	if err != nil {
		channel <- fmt.Sprintf("Link Not Working : %s, with reason %v", link, err.Error())
		return
	}

	channel <- fmt.Sprintf("Link Working : %s", link)
}
