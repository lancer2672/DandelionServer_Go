package middleware

import "net/http"

type MyHandlerFunc func(w http.ResponseWriter, r *http.Request) error

func errorHandler(h MyHandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
		}
	})
}
