package main

import (
	"log"

	"github.com/tinyrange/htm/v2"
	h "github.com/tinyrange/htm/v2/html"
	"github.com/tinyrange/htm/v2/htmx"
)

var containerId = h.NewId()

func Main() htm.Fragment {
	return h.Html(
		h.Head(
			h.Title("htmx demo"),
			htmx.Script,
		),
		h.Body(
			h.Div(containerId,
				h.Button(
					htmx.Post(h.Div(htm.Text("You clicked the button"))),
					htmx.Target(containerId),
					htm.Text("Press Me"),
				),
			),
		),
	)
}

func main() {
	log.Printf("Listening on: http://127.0.0.1:1512")
	err := htm.ListenAndServe("127.0.0.1:1512", Main())
	if err != nil {
		log.Fatal(err)
	}
}
