package htm

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type key int

var requestKey key

func RequestFromContext(ctx context.Context) (*http.Request, bool) {
	val, ok := ctx.Value(requestKey).(*http.Request)
	return val, ok
}

type routeMux struct {
	routes map[string]Fragment
}

func (mux *routeMux) Handle(w http.ResponseWriter, r *http.Request) (bool, error) {
	requestPath := r.URL.EscapedPath()

	route, ok := mux.routes[requestPath]
	if !ok {
		return false, nil
	}

	err := Render(r.Context(), w, route)
	if err != nil {
		return true, err
	}

	return true, err
}

var routeKey key

func RegisterRoute(ctx context.Context, url string, fragment Fragment) error {
	val, ok := ctx.Value(routeKey).(*routeMux)
	if !ok {
		return fmt.Errorf("could not get route mux")
	}

	val.routes[url] = fragment

	return nil
}

func ListenAndServe(addr string, f Fragment) error {
	return http.ListenAndServe(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		routeMux := &routeMux{
			routes: make(map[string]Fragment),
		}

		ctx := context.WithValue(r.Context(), requestKey, r)
		ctx = context.WithValue(ctx, routeKey, routeMux)

		err := WalkTree(ctx, f)
		if err != nil {
			log.Printf("failed to walk tree: %v", err)
			http.Error(w, "failed to walk tree", http.StatusInternalServerError)
			return
		}

		handled, err := routeMux.Handle(w, r)
		if err != nil {
			log.Printf("failed to handle special route: %v", err)
			http.Error(w, "failed to handle special route", http.StatusInternalServerError)
			return
		}

		if handled {
			return
		}

		err = Render(ctx, w, f)
		if err != nil {
			log.Printf("failed to render: %v", err)
			http.Error(w, "failed to render", http.StatusInternalServerError)
		}
	}))
}
