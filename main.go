package main

import (
	"net/http"
	"webby"
)

func main() {
	r := webby.Default()
	r.GET("/", func(c *webby.Context) {
		c.String(http.StatusOK, "Hello Kaustubh\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *webby.Context) {
		names := []string{"kaustubh"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
