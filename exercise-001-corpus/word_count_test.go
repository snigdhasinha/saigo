package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestCorpus(t *testing.T) {
  result := countWords("Are you serious? I knew you were.")
  assert.Equal(t, len(result), 6)
  assert.Equal(t, result[0].Word, "you")
  assert.Equal(t, result[0].Count, 2)
}