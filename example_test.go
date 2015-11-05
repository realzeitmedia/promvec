package promvec_test

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/realzeitmedia/promvec"
)

func Example() {
	var (
		machineName = "mach"
		labels      = [2]string{"method", "result"}
		methods     = []string{"GET", "PUT"}
		results     = []string{"success", "error"}
		labelValues [][2]string
	)

	for _, m := range methods {
		for _, r := range results {
			labelValues = append(labelValues, [2]string{m, r})
		}
	}

	promRequests := promvec.NewCounterVec2(
		prometheus.CounterOpts{
			Name:        "requests",
			Help:        "Requests",
			ConstLabels: prometheus.Labels{"machine": machineName},
		},
		labels,
		labelValues,
	)
	prometheus.MustRegister(promRequests)

	// ...

	promRequests.WithLabelValues([2]string{"GET", "error"}).Inc()
}
