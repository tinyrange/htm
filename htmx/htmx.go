package htmx

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/tinyrange/htm/v2"
	"github.com/tinyrange/htm/v2/html"
)

var Script = html.JavaScriptSrc("https://unpkg.com/htmx.org@1.9.2")

func MakeRandomUrl() string {
	return fmt.Sprintf("/api/u%s", strconv.FormatUint(rand.Uint64(), 36))
}

func Post(callback htm.Fragment) htm.Fragment {
	routeName := MakeRandomUrl()

	return htm.Group{
		htm.Route(routeName, callback),
		htm.Attr("hx-post", routeName),
	}
}

func Target(id html.Id) htm.Fragment {
	return htm.Attr("hx-target", "#"+string(id))
}
