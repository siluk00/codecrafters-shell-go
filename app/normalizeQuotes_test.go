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
		"'shell     hello' 'test''world' example''script",
	}

	expected := []string{
		"Hello you",
		"Hello you",
		"Hello you",
		"hi      you",
		"shell     hello testworld examplescript",
	}

	for i, _ := range strTest {
		acquired, _ := ParseQuotes(strTest[i])
		assert.Equal(t, expected[i], acquired)
	}
}
