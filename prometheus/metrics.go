// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package hangman

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Metrics represents Prometheus exposed metrics.
type Metrics struct {
	good, bad prometheus.Counter
	tally     prometheus.Gauge
}

// NewMetrics returns a new instance.
func NewMetrics() *Metrics {
	return &Metrics{
		good: promauto.NewCounter(prometheus.CounterOpts{
			Name: "hangman_good_guess_count",
			Help: "Counts number of good guesses",
			ConstLabels: map[string]string{
				"app":   "hangman",
				"guess": "good",
			},
		}),
		bad: promauto.NewCounter(prometheus.CounterOpts{
			Name: "hangman_bad_guess_count",
			Help: "Counts number of bad guesses",
			ConstLabels: map[string]string{
				"app":   "hangman",
				"guess": "bad",
			},
		}),
		tally: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "hangman_tally_total",
			Help: "The total number of game won or lost",
			ConstLabels: map[string]string{
				"app": "hangman",
			},
		}),
	}
}

// GoodGuess increments good case counter.
func (m *Metrics) GoodGuess() {
	m.good.Inc()
}

// BadGuess increments bad guess counter.
func (m *Metrics) BadGuess() {
	m.bad.Inc()
}

// UpdateTally updates tally.
func (m *Metrics) SetGameStatus(inc bool) {
	if inc {
		m.tally.Inc()
	} else {
		m.tally.Dec()
	}
}
