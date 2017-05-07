package helper

import (
	"log"
	"net/http"
	"time"
)

// Logger :http request logger
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// CORS: リモートアドレスからのアクセスを許可する
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
