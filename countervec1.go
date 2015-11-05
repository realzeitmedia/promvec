package promvec

import (
	"github.com/prometheus/client_golang/prometheus"
)

type CounterVec1 struct {
	desc     *prometheus.Desc
	counters map[string]prometheus.Counter
}

func NewCounterVec1(opts prometheus.CounterOpts, labelName string, labelValues []string) *CounterVec1 {
	v := CounterVec1{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(opts.Namespace, opts.Subsystem, opts.Name),
			opts.Help,
			[]string{labelName},
			opts.ConstLabels,
		),
		counters: map[string]prometheus.Counter{},
	}
	for _, lvs := range labelValues {
		labels := prometheus.Labels{}
		for cl, cv := range opts.ConstLabels {
			labels[cl] = cv
		}
		labels[labelName] = lvs
		opts.ConstLabels = labels
		v.counters[lvs] = prometheus.NewCounter(opts)
	}
	return &v
}

func (v *CounterVec1) Describe(ch chan<- *prometheus.Desc) {
	ch <- v.desc
}

func (v *CounterVec1) Collect(ch chan<- prometheus.Metric) {
	for _, c := range v.counters {
		ch <- c
	}
}

func (v *CounterVec1) WithLabelValues(lvs string) prometheus.Counter {
	if c, ok := v.counters[lvs]; !ok {
		panic("unexpected label values: " + lvs)
	} else {
		return c
	}
}
