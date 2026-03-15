package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeQuotes(t *testing.T) {
	strTest := []string{
		"'Hello you'",
		"'Hello' you",
		"Hello you",
		"'hi     ' you",
	}

	expected := []string{
		"Hello you",
		"Hello you",
		"Hello you",
		"hi      you",
	}

	for i, _ := range strTest {
		acquired, _ := NormalizeQuotes(strTest[i])
		assert.Equal(t, expected[i], acquired)
	}
}
