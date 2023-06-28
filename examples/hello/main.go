package main

import (
	"context"
	"log"
	"time"

	"github.com/tinyrange/htm/v2"
	h "github.com/tinyrange/htm/v2/html"
	tw "github.com/tinyrange/htm/v2/tailwind"
)

var CurrentTime = htm.Dynamic(func(ctx context.Context) ([]htm.Fragment, error) {
	return []htm.Fragment{h.Textf("%s", time.Now().String())}, nil
})

func main() {
	log.Printf("Listening on: http://127.0.0.1:1512")
	err := htm.ListenAndServe("127.0.0.1:1512", h.Html(
		h.Head(
			h.Title("Hello, World"),
			h.JavaScriptSrc("https://cdn.tailwindcss.com"),
			h.JavaScriptSrc("https://unpkg.com/htmx.org@1.9.2"),
		),
		h.Body(
			h.Div(tw.Container,
				h.Span(tw.Font.Weight.Bold, h.Text("The current time is: ")),
				CurrentTime,
			),
		),
	))
	if err != nil {
		log.Fatal(err)
	}
}
