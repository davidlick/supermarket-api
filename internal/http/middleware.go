package http

import "net/http"

// setResponseHeaders is a middleware that accepts a map[string]string of headers:values and sets them for all responses.
func setResponseHeaders(headers map[string]string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for header, value := range headers {
				w.Header().Set(header, value)
			}

			next.ServeHTTP(w, r)
		})
	}
}
