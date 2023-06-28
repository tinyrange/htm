package htm

import (
	"log"
	"net/http"
)

func ListenAndServe(addr string, f Fragment) error {
	return http.ListenAndServe(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := Render(r.Context(), w, f)
		if err != nil {
			log.Printf("failed to render: %v", err)
		}
	}))
}
