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
	<<!!YOUR_CODE!!>> Declare your prom metrics to track good/bad guesses and tally metrics
}

// NewMetrics returns a new instance.
func NewMetrics() *Metrics {
	return &Metrics{
		<<!!YOUR_CODE!!>> Initialize prom counters
	}
}

// GoodGuess increments good case counter.
func (m *Metrics) GoodGuess() {
	<<!!YOUR_CODE!!>> Track your good guesses metrics
}

// BadGuess increments bad guess counter.
func (m *Metrics) BadGuess() {
	<<!!YOUR_CODE!!>> Track your bad guesses metrics
}

// UpdateTally updates tally.
func (m *Metrics) SetGameStatus(won bool) {
	<<!!YOUR_CODE!!>> Track your tally metrics if game is won or lost
}
