package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var requestCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	},
	[]string{"method", "handler"},
)

func init() {
	prometheus.MustRegister(requestCount)
}

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount.With(prometheus.Labels{"method": r.Method, "handler": "/"}).Inc()
		next.ServeHTTP(w, r)
	})
}
