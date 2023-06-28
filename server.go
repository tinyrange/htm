package htm

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type key int

var requestKey key

type requestContext struct {
	request *http.Request
	routes  map[string]Fragment
	dynamic bool
}

func MakeDynamicContext(ctx context.Context) (context.Context, error) {
	val, ok := ctx.Value(requestKey).(*requestContext)
	if !ok {
		return nil, fmt.Errorf("could not get current request context")
	}

	return context.WithValue(ctx, requestKey, &requestContext{
		request: val.request,
		routes:  nil,
		dynamic: true,
	}), nil
}

func RequestFromContext(ctx context.Context) (*http.Request, bool) {
	val, ok := ctx.Value(requestKey).(*requestContext)
	if !ok {
		return nil, false
	}
	return val.request, true
}

func (mux *requestContext) Handle(ctx context.Context, w http.ResponseWriter, r *http.Request) (bool, error) {
	requestPath := r.URL.EscapedPath()

	route, ok := mux.routes[requestPath]
	if !ok {
		return false, nil
	}

	dynCtx, err := MakeDynamicContext(ctx)
	if err != nil {
		return false, nil
	}

	err = Render(dynCtx, w, route)
	if err != nil {
		return true, err
	}

	return true, err
}

type route struct {
	url      string
	fragment Fragment
}

func (r *route) Children(ctx context.Context) ([]Fragment, error) {
	val, ok := ctx.Value(requestKey).(*requestContext)
	if !ok {
		return nil, fmt.Errorf("could not get request context")
	}

	val.routes[r.url] = r.fragment

	return []Fragment{}, nil
}

func (*route) Render(ctx context.Context, parent Node) error {
	// Routes are invisible.
	return nil
}

var (
	_ Fragment = &route{}
)

func Route(url string, f Fragment) Fragment {
	return &route{url: url, fragment: f}
}

func ListenAndServe(addr string, f Fragment) error {
	return http.ListenAndServe(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqCtx := &requestContext{
			request: r,
			routes:  make(map[string]Fragment),
		}

		ctx := context.WithValue(r.Context(), requestKey, reqCtx)

		err := WalkTree(ctx, f)
		if err != nil {
			log.Printf("failed to walk tree: %v", err)
			http.Error(w, "failed to walk tree", http.StatusInternalServerError)
			return
		}

		handled, err := reqCtx.Handle(ctx, w, r)
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
