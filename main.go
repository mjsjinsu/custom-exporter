package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	log "github.com/sirupsen/logrus"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	log.Info("Start to serve on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}