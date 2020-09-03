package main

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/mjsjinsu/custom-exporter/collector"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/promlog"
	"github.com/prometheus/common/promlog/flag"
	"github.com/prometheus/common/version"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"net/http"
	"os"
	"sort"
)

var (
	listenAddress = kingpin.Flag(
		"web.listen-address",
		"Address on which to expose metrics and web interface.",
	).Default(":8080").String()

	metricsPath = kingpin.Flag(
		"web.telemetry-path",
		"Path under which to expose metrics.",
	).Default("/metrics").String()

	disableDefaultCollectors = kingpin.Flag(
		"collector.disable-defaults",
		"Set all collectors to disabled by default.",
	).Default("false").Bool()

	disableExporterMetrics = kingpin.Flag(
		"web.disable-exporter-metrics",
		"Exclude metrics about the exporter itself (promhttp_*, process_*, go_*).",
	).Bool()
)

type handler struct {
	unfilteredHandler      http.Handler
	exporterMetricRegistry *prometheus.Registry
	includeExporterMetrics bool
	logger                 log.Logger
}

func newHandler(includeExporterMetrics bool, logger log.Logger) *handler {
	h := &handler{
		exporterMetricRegistry: prometheus.NewRegistry(),
		includeExporterMetrics: includeExporterMetrics,
		logger:                 logger,
	}
	if h.includeExporterMetrics {
		h.exporterMetricRegistry.MustRegister(
			prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}),
			prometheus.NewGoCollector(),
		)
	}
	if innerHandler, err := h.innerHandler(); err != nil {
		panic(fmt.Sprintf("Couldn't create metrics handler: %s", err))
	} else {
		h.unfilteredHandler = innerHandler
	}
	return h
}

// http.Handler Interface의 ServeHTTP 메소드 구현체
func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filters := r.URL.Query()["collect[]"]
	level.Debug(h.logger).Log("msg", "collect query:", "filters", filters)

	if len(filters) == 0 {
		// No filters, use the prepared unfiltered handler.
		h.unfilteredHandler.ServeHTTP(w, r)
		return
	}

	filteredHandler, err := h.innerHandler(filters...)
	if err != nil {
		level.Warn(h.logger).Log("msg", "Couldn't create filtered metrics handler:", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Couldn't create filtered metrics handler: %s", err)))
		return
	}
	filteredHandler.ServeHTTP(w, r)
}

func (h *handler) innerHandler(filters ...string) (http.Handler, error) {
	cc, err := collector.NewCustomCollector(h.logger, filters...)
	if err != nil {
		return nil, fmt.Errorf("couldn't create collector: %s", err)
	}

	if len(filters) == 0 {
		level.Info(h.logger).Log("msg", "Enabled collectors")
		collectors := []string{}
		for n := range cc.Collectors {
			collectors = append(collectors, n)
		}
		sort.Strings(collectors)
		for _, c := range collectors {
			level.Info(h.logger).Log("collector", c)
		}
	}

	r := prometheus.NewRegistry()
	//버전 정보 등록
	r.MustRegister(version.NewCollector("custom_exporter"))
	if err := r.Register(cc); err != nil {
		return nil, fmt.Errorf("couldn't register custom collector: %s", err)
	}

	handler := promhttp.HandlerFor(prometheus.Gatherers{h.exporterMetricRegistry, r},
		promhttp.HandlerOpts{
			ErrorLog:            nil,
			ErrorHandling:       promhttp.ContinueOnError,
			Registry:            h.exporterMetricRegistry,
			DisableCompression:  false,
			MaxRequestsInFlight: 0,
			Timeout:             0,
			EnableOpenMetrics:   false,
		})

	if h.includeExporterMetrics {
		handler = promhttp.InstrumentMetricHandler(h.exporterMetricRegistry, handler)
	}

	return handler, nil
}

func main() {
	//promlog 설정
	promlogConfig := &promlog.Config{}

	//커맨드 라인 인수 설정
	flag.AddFlags(kingpin.CommandLine, promlogConfig)

	kingpin.Version(version.Print("custom_exporter"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	//기본 로그 출력
	logger := promlog.New(promlogConfig)

	if *disableDefaultCollectors {
		collector.DisableDefaultCollectors()
	}

	level.Info(logger).Log("msg", "Starting custom_exporter", "version", version.Info())
	level.Info(logger).Log("msg", "Build context", "build_context", version.BuildContext())

	//http.Handle(*metricsPath, promhttp.Handler())
	http.Handle(*metricsPath, newHandler(*disableExporterMetrics, logger))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>Custom Exporter</title></head>
			<body>
			<h1>Custom Exporter</h1>
			<p><a href="` + *metricsPath + `">Metrics</a></p>
			</body>
			</html>`))
	})
	level.Info(logger).Log("msg", "Listening on", "address", *listenAddress)

	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		level.Error(logger).Log("msg", "Error starting HTTP server", "err", err)
		os.Exit(1)
	}
}
