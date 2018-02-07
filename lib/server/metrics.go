package server

import (
	"github.com/prometheus/client_golang/prometheus"
)

var metricsNamespace = "centrifugo"
var metricsSubsystem = "httpserver"

var (
	apiHandlerDurationSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Namespace:  metricsNamespace,
		Subsystem:  metricsSubsystem,
		Name:       "api_request_duration_seconds",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		Help:       "TODO.",
	})

	apiCommandDurationSummary = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace:  metricsNamespace,
		Subsystem:  metricsSubsystem,
		Name:       "api_request_command_duration_seconds",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		Help:       "TODO.",
	}, []string{"method"})

	transportConnectCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: metricsNamespace,
		Subsystem: metricsSubsystem,
		Name:      "transport_connect_count",
		Help:      "TODO.",
	}, []string{"transport"})

	transportMessagesSent = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: metricsNamespace,
		Subsystem: metricsSubsystem,
		Name:      "transport_messages_sent",
		Help:      "TODO.",
	}, []string{"transport"})

	transportBytesOut = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: metricsNamespace,
		Subsystem: metricsSubsystem,
		Name:      "transport_bytes_out",
		Help:      "TODO.",
	}, []string{"transport"})
)

func init() {
	prometheus.MustRegister(apiHandlerDurationSummary)
	prometheus.MustRegister(apiCommandDurationSummary)
	prometheus.MustRegister(transportConnectCount)
}
