package promvec

import (
	"github.com/prometheus/client_golang/prometheus"
)

type GaugeVec1 struct {
	desc   *prometheus.Desc
	gauges map[string]prometheus.Gauge
}

func NewGaugeVec1(opts prometheus.GaugeOpts, labelName string, labelValues []string) *GaugeVec1 {
	v := GaugeVec1{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(opts.Namespace, opts.Subsystem, opts.Name),
			opts.Help,
			[]string{labelName},
			opts.ConstLabels,
		),
		gauges: map[string]prometheus.Gauge{},
	}
	for _, lvs := range labelValues {
		labels := prometheus.Labels{}
		for cl, cv := range opts.ConstLabels {
			labels[cl] = cv
		}
		labels[labelName] = lvs
		opts.ConstLabels = labels
		v.gauges[lvs] = prometheus.NewGauge(opts)
	}
	return &v
}

func (v *GaugeVec1) Describe(ch chan<- *prometheus.Desc) {
	ch <- v.desc
}

func (v *GaugeVec1) Collect(ch chan<- prometheus.Metric) {
	for _, c := range v.gauges {
		ch <- c
	}
}

func (v *GaugeVec1) WithLabelValues(lvs string) prometheus.Gauge {
	if c, ok := v.gauges[lvs]; !ok {
		panic("unexpected label values: " + lvs)
	} else {
		return c
	}
}
