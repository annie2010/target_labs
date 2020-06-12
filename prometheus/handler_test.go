// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package hangman_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gopherland/target_labs/prometheus/hangman"
	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	m := new(metrics)
	h, err := hangman.NewHandler("testdata/one.txt", m)
	assert.Nil(t, err)

	var (
		rr   = httptest.NewRecorder()
		r, _ = http.NewRequest("GET", "http://example.com/api/v1/new_game", nil)
	)

	h.NewGame(rr, r)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGuess(t *testing.T) {
	uu := map[string]struct {
		guesses          []rune
		guess            rune
		good, bad, tally int
	}{
		"good": {guess: 'f', guesses: []rune{'r', 'e', 'd'}, good: 1, tally: 1},
		"bad":  {guess: 'z', guesses: []rune{'r', 'e', 'd'}, bad: 1, tally: 0},
		"lost": {guess: 'z', guesses: []rune{'a', 'b', 'c', 'g', 'h', 'i'}, bad: 1, tally: -1},
	}

	for k := range uu {
		u := uu[k]

		m := new(metrics)
		h, err := hangman.NewHandler("testdata/one.txt", m)
		assert.Nil(t, err)
		body := struct {
			Game  *hangman.Game
			Guess rune
		}{
			Game:  hangman.NewGame("fred"),
			Guess: u.guess,
		}
		for _, r := range u.guesses {
			body.Game.Guess(r)
		}

		t.Run(k, func(t *testing.T) {
			raw, err := json.Marshal(&body)
			assert.Nil(t, err)

			var (
				rr   = httptest.NewRecorder()
				r, _ = http.NewRequest("GET", "http://example.com/api/v1/guess", bytes.NewBuffer(raw))
			)

			h.Guess(rr, r)
			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Equal(t, u.good, m.good)
			assert.Equal(t, u.bad, m.bad)
			assert.Equal(t, u.tally, m.status)
		})
	}
}

type metrics struct {
	good, bad, status int
}

func (m *metrics) SetGameStatus(won bool) {
	if won {
		m.status++
	} else {
		m.status--
	}
}

func (m *metrics) GoodGuess() {
	m.good++

}
func (m *metrics) BadGuess() {
	m.bad++
}
