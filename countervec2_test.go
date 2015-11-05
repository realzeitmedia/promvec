package promvec

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

func TestCounterVec2Empty(t *testing.T) {
	var (
		desc = make(chan *prometheus.Desc, 10)
		coll = make(chan prometheus.Metric, 10)
	)
	v := NewCounterVec2(
		prometheus.CounterOpts{
			Name:        "test",
			Help:        "test",
			ConstLabels: nil,
		},
		[2]string{},
		[][2]string{},
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

func TestCounterVec2(t *testing.T) {
	var (
		desc = make(chan *prometheus.Desc, 10)
		coll = make(chan prometheus.Metric, 10)
	)
	v := NewCounterVec2(
		prometheus.CounterOpts{
			Name:        "test",
			Help:        "test",
			ConstLabels: nil,
		},
		[2]string{"first", "second"},
		[][2]string{
			{"a", "b"},
			{"aa", "bb"},
			{"aa", "cc"},
		},
	)

	v.Describe(desc)
	d, ok := <-desc
	if !ok || d == nil || len(d.String()) == 0 {
		t.Errorf("no description")
	}

	v.Collect(coll)
	if have, want := len(coll), 3; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	c := v.WithLabelValues([2]string{"aa", "bb"})
	d = c.Desc()
	if d == nil || len(d.String()) == 0 {
		t.Errorf("no counter description")
	}
}
