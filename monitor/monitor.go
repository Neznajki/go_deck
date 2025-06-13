package main

import (
	"fmt"
	"net/http"
	"time"
)

type RequestData struct {
	Link      string
	c         chan RequestData
	terminate bool
	tries     int
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
