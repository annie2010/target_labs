// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package hangman

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

// Score tracks game metrics.
type Scorer interface {
	// GameStatus sets the current game status.
	SetGameStatus(bool)

	// GoodGuess increments good guesses.
	GoodGuess()

	// BadGuess increments bad guesses.
	BadGuess()
}

var _ Scorer = (*Metrics)(nil)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Handler struct {
	dictionary *Dictionary
	scorer     Scorer
}

func NewHandler(file string, scorer Scorer) (*Handler, error) {
	d, err := NewDictionary(file)
	if err != nil {
		return nil, err
	}

	return &Handler{
		dictionary: d,
		scorer:     scorer,
	}, nil
}

// Guess checks if guess if valid or not.
func (h *Handler) Guess(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Game  *Game `json:"game"`
		Guess rune  `json:"guess"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	outcome := req.Game.Guess(rune(req.Guess))
	switch outcome {
	case GoodGuess:
		h.scorer.GoodGuess()
	case BadGuess:
		h.scorer.BadGuess()
	}

	switch req.Game.Tally.Status {
	case Won:
		h.scorer.SetGameStatus(true)
	case Lost:
		h.scorer.SetGameStatus(false)
	}

	raw, err := json.Marshal(req.Game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if _, err := w.Write(raw); err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
	}
}

// NewGame starts a new game.
func (h *Handler) NewGame(w http.ResponseWriter, r *http.Request) {
	g := NewGame(h.pick())

	buff, err := json.Marshal(g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if _, err := w.Write(buff); err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
	}
}

func (h *Handler) pick() string {
	words := h.dictionary.Words()
	return words[rand.Intn(len(words))]
}
