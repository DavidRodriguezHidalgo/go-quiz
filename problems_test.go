package main

import (
	"testing"

	"./problems"
	"github.com/stretchr/testify/assert"
)

func TestReadProblems(t *testing.T) {
  p := problems.ReadProblems("problems.csv")
  assert.Equal(t, len(p), 12, "It should read the 12 problems")
}

