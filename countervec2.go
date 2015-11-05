package promvec

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

// CounterVec2 bundles counters that have different labels in two dimensions,
// with a fixed set of possible label combinations.
type CounterVec2 struct {
	desc     *prometheus.Desc
	counters map[[2]string]prometheus.Counter
}

// NewCounterVec2 creates a CounterVec2, given counter options and definitions
// of label names and label values.
func NewCounterVec2(opts prometheus.CounterOpts, labelNames [2]string, labelValues [][2]string) *CounterVec2 {
	v := CounterVec2{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(opts.Namespace, opts.Subsystem, opts.Name),
			opts.Help,
			labelNames[:],
			opts.ConstLabels,
		),
		counters: map[[2]string]prometheus.Counter{},
	}
	for _, lvs := range labelValues {
		labels := prometheus.Labels{}
		for cl, cv := range opts.ConstLabels {
			labels[cl] = cv
		}
		for i, l := range labelNames {
			labels[l] = lvs[i]
		}
		opts.ConstLabels = labels
		v.counters[lvs] = prometheus.NewCounter(opts)
	}
	return &v
}

// Describe implements Collector.
func (v *CounterVec2) Describe(ch chan<- *prometheus.Desc) {
	ch <- v.desc
}

// Collect implements Collector.
func (v *CounterVec2) Collect(ch chan<- prometheus.Metric) {
	for _, c := range v.counters {
		ch <- c
	}
}

// WithLabelValues returns the Counter for the given label values, and
// panics if the labels are invalid.
func (v *CounterVec2) WithLabelValues(lvs [2]string) prometheus.Counter {
	if c, ok := v.counters[lvs]; !ok {
		panic("unexpected label values: " + strings.Join(lvs[:], ", "))
	} else {
		return c
	}
}
