package htm

import (
	"context"
	"log"
	"net/http"
)

type key int

var requestKey key

func RequestFromContext(ctx context.Context) (*http.Request, bool) {
	val, ok := ctx.Value(requestKey).(*http.Request)
	return val, ok
}

func ListenAndServe(addr string, f Fragment) error {
	return http.ListenAndServe(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), requestKey, r)

		err := Render(ctx, w, f)
		if err != nil {
			log.Printf("failed to render: %v", err)
		}
	}))
}
