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
					if strings.TrimSpace(outsideContent.String()) != "" {
						content.WriteString(strings.Join(strings.Fields(outsideContent.String()), " "))
					} else if content.String() != "" {
						content.WriteRune(' ')
					}
					outsideContent.Reset()
					flipped = false
				}
				content.WriteRune(r)
			} else {
				outsideContent.WriteRune(r)
				flipped = false
			}
		}
	}

	if counter%2 != 0 {
		return "", errors.New("Bad argument")
	}

	if !inQuote {
		content.WriteString(strings.Join(strings.Fields(outsideContent.String()), " "))
	}

	return content.String(), nil
}
