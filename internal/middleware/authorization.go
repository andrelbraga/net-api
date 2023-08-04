package middleware

import (
	"net/http"
)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if header := r.Header.Get("Auth-Key"); header == "" {
			w.WriteHeader(401)
			w.Write([]byte("401 Unauthorized!"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
