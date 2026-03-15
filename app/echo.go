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

// NormalizeQuotes is a function for echo to work with 'literals    ' input
func NormalizeQuotes(args string) (string, error) {
	content := strings.Builder{}
	inQuote := false
	flipped := false
	outsideContent := strings.Builder{}
	counter := 0
	args = strings.TrimLeft(args, " ")
	args = strings.TrimRight(args, " ")

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
					outsideString := outsideContent.String()
					if strings.TrimSpace(outsideString) != "" {
						if outsideString[0] == ' ' {
							content.WriteRune(' ')
						}
						content.WriteString(strings.Join(strings.Fields(outsideString), " "))
						if outsideString[len(outsideString)-1] == ' ' {
							content.WriteRune(' ')
						}
					} else if outsideString != "" {
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
