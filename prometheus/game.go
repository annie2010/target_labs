// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package hangman

type Outcome int

const (
	DuplicateGuess Outcome = 1 << iota
	GoodGuess
	BadGuess
	GameOver
)

// Game a hangman game
type Game struct {
	Word    string `json:"word"`
	Guesses []rune `json:"guesses"`
	Tally   *Tally `json:"tally"`
}

// NewGame initializes a hangman game
func NewGame(word string) *Game {
	return &Game{Word: word, Tally: NewTally([]rune(word)), Guesses: []rune{}}
}

// Guess a new letter
func (g *Game) Guess(guess rune) Outcome {
	return g.validateGuess(guess)
}

func (g *Game) validateGuess(guess rune) Outcome {
	if g.Tally.Status == Won || g.Tally.Status == Lost {
		return GameOver
	}

	if g.alreadyGuessed(guess) {
		g.Tally.Status = AlreadyGuessed
		return DuplicateGuess
	}

	g.Guesses = append(g.Guesses, guess)
	defer g.Tally.Update([]rune(g.Word), g.Guesses)
	if g.inWord(guess) {
		return GoodGuess
	}
	g.Tally.TurnsLeft--

	return BadGuess
}

func (g *Game) alreadyGuessed(guess rune) bool {
	for _, l := range g.Guesses {
		if l == guess {
			return true
		}
	}
	return false
}

func (g *Game) inWord(guess rune) bool {
	for _, l := range g.Word {
		if l == guess {
			return true
		}
	}
	return false
}
