package htmx

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/tinyrange/htm/v2"
	"github.com/tinyrange/htm/v2/html"
)

var Script = html.JavaScriptSrc("https://unpkg.com/htmx.org@1.9.2")

func MakeRandomUrl() string {
	return fmt.Sprintf("/dynamic/u%s", strconv.FormatUint(rand.Uint64(), 36))
}

func Post(callback htm.Fragment) htm.Fragment {
	routeName := MakeRandomUrl()

	return htm.Dynamic(func(ctx context.Context) ([]htm.Fragment, error) {
		err := htm.RegisterRoute(ctx, routeName, callback)
		if err != nil {
			return nil, fmt.Errorf("failed to register route: %v", err)
		}

		return []htm.Fragment{htm.Attr("hx-post", routeName)}, nil
	})
}

func Target(id html.Id) htm.Fragment {
	return htm.Attr("hx-target", "#"+string(id))
}
