package main

import (
	"gotest.tools/v3/assert"
	"testing"
	"time"
)

func Test_executeWithRetry_stop(t *testing.T) {
	handleTimeout(10 * time.Second)
	rd := RequestData{Link: "https://www.google.com", tries: 10, c: make(chan RequestData)}

	startTime := time.Now().Local()
	go executeWithRetry(rd)

	newRd := <-rd.c
	assert.Equal(t, newRd.terminate, true)
	assert.Equal(t, newRd.tries, 10)
	assert.Equal(t, rd.terminate, false)

	assert.Equal(t, true, time.Since(startTime) < time.Second)
}

func Test_executeWithRetry_lastTry(t *testing.T) {
	handleTimeout(10 * time.Second)
	rd := RequestData{Link: "https://www.google.com", tries: 9, c: make(chan RequestData)}

	startTime := time.Now().Local()
	go executeWithRetry(rd)

	newRd := <-rd.c
	assert.Equal(t, newRd.terminate, false)
	assert.Equal(t, newRd.tries, 10)
	assert.Equal(t, rd.terminate, false)

	assert.Equal(t, true, time.Since(startTime) > time.Second)
	assert.Equal(t, true, time.Since(startTime) < 3*time.Second)
}

func Test_executeWithRetry_wrongLink(t *testing.T) {
	handleTimeout(10 * time.Second)
	rd := RequestData{Link: "https:www.google.com", c: make(chan RequestData)}

	startTime := time.Now().Local()
	go executeWithRetry(rd)

	newRd := <-rd.c
	assert.Equal(t, newRd.terminate, true)
	assert.Equal(t, newRd.tries, 1)
	assert.Equal(t, rd.terminate, false)

	assert.Equal(t, true, time.Since(startTime) < time.Second)
}

func handleTimeout(t time.Duration) {
	go func() {
		time.Sleep(t)
		panic("wait timeout reached")
	}()
}
