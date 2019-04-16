package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {

	counterTest := prometheus.NewCounter(prometheus.CounterOpts{
		Name:        "throughput",
		Help:        "api throughput",
		ConstLabels: prometheus.Labels{"type": "in"},
	})
	prometheus.Register(counterTest)

	handler := promhttp.Handler()
	http.ListenAndServe(":8787", handler)
}
