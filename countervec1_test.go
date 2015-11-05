package promvec

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

func TestCounterVec1Empty(t *testing.T) {
	var (
		desc = make(chan *prometheus.Desc, 10)
		coll = make(chan prometheus.Metric, 10)
	)
	v := NewCounterVec1(
		prometheus.CounterOpts{
			Name:        "test",
			Help:        "test",
			ConstLabels: nil,
		},
		"",
		[]string{},
	)

	v.Describe(desc)
	d, ok := <-desc
	if !ok || d == nil || len(d.String()) == 0 {
		t.Errorf("no description")
	}

	v.Collect(coll)
	if have, want := len(coll), 0; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
}

func TestCounterVec1(t *testing.T) {
	var (
		desc = make(chan *prometheus.Desc, 10)
		coll = make(chan prometheus.Metric, 10)
	)
	v := NewCounterVec1(
		prometheus.CounterOpts{
			Name:        "test",
			Help:        "test",
			ConstLabels: nil,
		},
		"first",
		[]string{
			"a",
			"aa",
		},
	)

	v.Describe(desc)
	d, ok := <-desc
	if !ok || d == nil || len(d.String()) == 0 {
		t.Errorf("no description")
	}

	v.Collect(coll)
	if have, want := len(coll), 2; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	c := v.WithLabelValues("aa")
	d = c.Desc()
	if d == nil || len(d.String()) == 0 {
		t.Errorf("no counter description")
	}
}
