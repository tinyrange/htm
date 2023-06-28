package main

import (
	"log"

	"github.com/tinyrange/htm/v2"
	h "github.com/tinyrange/htm/v2/html"
)

func main() {
	log.Printf("Listening on: http://127.0.0.1:1512")
	err := htm.ListenAndServe("127.0.0.1:1512", h.Html(
		h.Head(
			h.Title("Hello, World"),
			h.LinkCSS("https://cdn.tailwindcss.com"),
			h.JavaScriptSrc("https://unpkg.com/htmx.org@1.9.2"),
		),
		h.Body(
			h.Div(h.Url()),
		),
	))
	if err != nil {
		log.Fatal(err)
	}
}
