package main

import (
	"errors"
	"fmt"
	"strings"
)

func echoCommand(args string) {
	phrase, err := NormalizeQuotes(args)
	if err != nil {
		return
	}
	fmt.Printf("%s\n", phrase)
}

func NormalizeQuotes(args string) (string, error) {
	content := strings.Builder{}
	inQuote := false
	flipped := false
	outsideContent := strings.Builder{}
	counter := 0

	for _, r := range args {
		if r == '\'' {
			counter++

			if inQuote {
				inQuote = false
				flipped = true
			} else {
				inQuote = true
				flipped = true
			}
		} else {
			if inQuote {
				if flipped {
					content.WriteString(strings.Join(strings.Fields(outsideContent.String()), " "))
					outsideContent.Reset()
					flipped = false
				}
				content.WriteRune(r)
			} else {
				outsideContent.WriteRune(r)
			}
		}
	}

	if counter%2 != 0 {
		return "", errors.New("Bad argument")
	}

	if !inQuote {
		content.WriteString(outsideContent.String())
	}

	return content.String(), nil
}
