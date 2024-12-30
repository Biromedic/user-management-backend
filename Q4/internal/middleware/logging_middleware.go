package middleware

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logrus.Infof("Started %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		logrus.Infof("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	})
}
