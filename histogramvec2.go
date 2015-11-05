package promvec

import (
	"github.com/prometheus/client_golang/prometheus"
)

// HistogramVec2 bundles histogramss that have different labels in two dimensions,
// with a fixed set of possible label combinations.
type HistogramVec2 struct {
	desc       *prometheus.Desc
	histograms map[[2]string]prometheus.Histogram
}

// NewHistogramVec2 creates a HistogramVec2, given histogram options and definitions
// of label names and label values.
func NewHistogramVec2(opts prometheus.HistogramOpts, labelNames [2]string, labelValues [][2]string) *HistogramVec2 {
	v := HistogramVec2{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(opts.Namespace, opts.Subsystem, opts.Name),
			opts.Help,
			labelNames[:],
			opts.ConstLabels,
		),
		histograms: map[[2]string]prometheus.Histogram{},
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
		v.histograms[lvs] = prometheus.NewHistogram(opts)
	}
	return &v
}

// Describe implements Collector.
func (v *HistogramVec2) Describe(ch chan<- *prometheus.Desc) {
	ch <- v.desc
}

// Collect implements Collector.
func (v *HistogramVec2) Collect(ch chan<- prometheus.Metric) {
	for _, c := range v.histograms {
		ch <- c
	}
}

// WithLabelValues returns the Histogram for the given label values, and
// panics if the labels are invalid.
func (v *HistogramVec2) WithLabelValues(lvs [2]string) prometheus.Histogram {
	if c, ok := v.histograms[lvs]; !ok {
		panic("unexpected label values: " + lvs[0] + ", " + lvs[1])
	} else {
		return c
	}
}
