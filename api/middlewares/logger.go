package middlewares

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler, name string, logging *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		defer logging.Printf("%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(startTime),
		)
		next.ServeHTTP(w, r)
	})
}
