package main

import (
	"log"
	"lulu"
	"net/http"
	"time"
)

func onlyForV2() lulu.HandlerFunc {
	return func(c *lulu.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := lulu.Default()
	r.GET("/", func(c *lulu.Context) {
		c.String(http.StatusOK, "Hello luluktutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *lulu.Context) {
		names := []string{"luluktutu"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
