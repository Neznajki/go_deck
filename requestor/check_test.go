package main

import (
	"gotest.tools/v3/assert"
	"testing"
	"time"
)

func Test_CheckSite(t *testing.T) {
	c := make(chan string)
	go func() {
		time.Sleep(10 * time.Second)
		c <- "wait timeout reached"
	}()
	startTime := time.Now().Local()
	go checkSite("https://www.google.com", c)

	assert.Equal(t, <-c, "Link Working : https://www.google.com")

	go checkSite("https://www.this.will.not.work", c)

	assert.Equal(t, <-c, "Link Not Working : https://www.this.will.not.work, with reason Get \"https://www.this.will.not.work\": remote error: tls: handshake failure")

	assert.Equal(t, true, time.Since(startTime) < time.Second)
}
