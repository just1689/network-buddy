package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	HTTPCalls = promauto.NewCounter(prometheus.CounterOpts{
		Name: "vote_http_calls",
		Help: "The total number of http calls",
	})

	Sessions = promauto.NewCounter(prometheus.CounterOpts{
		Name: "vote_sessions",
		Help: "The total number of sessions",
	})

	Votes = promauto.NewCounter(prometheus.CounterOpts{
		Name: "vote_votes",
		Help: "The total number of votes",
	})

	A = promauto.NewCounter(prometheus.CounterOpts{
		Name: "vote_A",
		Help: "The total number of A",
	})

	B = promauto.NewCounter(prometheus.CounterOpts{
		Name: "vote_B",
		Help: "The total number of B",
	})

	Active = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "vote_active_connections",
		Help: "The actuall number currently connected",
	})
)
