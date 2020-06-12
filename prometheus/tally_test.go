// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package hangman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateLettersMiss(t *testing.T) {
	res := updateLetters([]rune("hello"), []rune("x"))

	assert.Equal(t, "_____", string(res))
}

func TestUpdateLettersHit(t *testing.T) {
	res := updateLetters([]rune("hello"), []rune("l"))

	assert.Equal(t, "__ll_", string(res))
}
