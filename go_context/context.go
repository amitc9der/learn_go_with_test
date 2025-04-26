package gocontext

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

// managing long runing process
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data, err := store.Fetch(ctx)
		if err != nil {
			log.Printf("Error occured %q", err)
			return //
		}
		fmt.Fprint(w, data)
	}
}
