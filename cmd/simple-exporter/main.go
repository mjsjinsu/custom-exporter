package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var (
	addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
)

var (
	//Define the metrics we wish to expose
	fooMetric = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace:   "",
		Subsystem:   "",
		Name:        "foo_metric",
		Help:        "Shows whether a foo has occurred in our cluster",
		ConstLabels: nil,
	})
	barMetric = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace:   "",
		Subsystem:   "",
		Name:        "bar_metric",
		Help:        "Shows whether a bar has occurred in our cluster",
		ConstLabels: nil,
	})
)

func init() {
	// Register the Gauge and the Counter with Prometheus's default registry.
	prometheus.MustRegister(fooMetric)
	prometheus.MustRegister(barMetric)

	fooMetric.Set(0)
	barMetric.Inc()
}

func main() {
	flag.Parse()
	http.Handle("/metrics", promhttp.Handler())
	log.Info("Start to serve on port" + *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
