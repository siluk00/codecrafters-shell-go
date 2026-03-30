package main

import (
	"errors"
	"strings"
)

// NormalizeQuotes is a function for echo to work with 'literals    ' input
func ParseQuotes(args string) (string, error) {
	content := strings.Builder{}
	inQuote := false
	outsideContent := strings.Builder{}
	counter := 0
	args = strings.TrimLeft(args, " ")
	args = strings.TrimRight(args, " ")

	for _, r := range args {
		if r == '\'' {
			counter++

			if inQuote {
				inQuote = false
			} else {
				inQuote = true

				outsideString := outsideContent.String()

				if strings.TrimSpace(outsideString) != "" {
					if strings.HasPrefix(outsideString, " ") {
						content.WriteRune(' ')
					}
					content.WriteString(strings.Join(strings.Fields(outsideString), " "))
					if strings.HasSuffix(outsideString, " ") {
						content.WriteRune(' ')
					}
				} else if outsideString != "" {
					content.WriteRune(' ')
				}
				outsideContent.Reset()
			}
		} else {
			if inQuote {
				content.WriteRune(r)
			} else {
				outsideContent.WriteRune(r)
			}
		}
	}

	if counter%2 != 0 {
		return "", errors.New("Bad argument: unbalanced quotes")
	}

	if !inQuote {
		content.WriteString(strings.Join(strings.Fields(outsideContent.String()), " "))
	}

	return content.String(), nil
}
