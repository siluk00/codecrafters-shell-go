package main

import (
	"errors"
	"strings"
)

// NormalizeQuotes is a function for echo to work with 'literals    ' input
func tokenizer(args string) ([]string, error) {
	insideContent := strings.Builder{}
	content := []string{}
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
				insideString := insideContent.String()
				if outsideString == "" {
					continue
				} else if strings.Trim(outsideString, " ") == "" {
					if insideString != "" {
						content = append(content, insideString)
					}
					outsideContent.Reset()
					insideContent.Reset()
				} else {
					if strings.HasPrefix(outsideString, " ") && strings.HasSuffix(outsideString, " ") {
						if insideString != "" {
							content = append(content, insideString)
						}
						content = append(content, strings.Fields(outsideString)...)
						outsideContent.Reset()
						insideContent.Reset()
					} else if strings.HasPrefix(outsideString, " ") {
						if insideString != "" {
							content = append(content, insideString)
						}
						fields := strings.Fields(outsideString)
						content = append(content, fields[0:len(fields)-1]...)
						insideContent.WriteString(fields[0])
						outsideContent.Reset()
					} else if strings.HasSuffix(outsideString, " ") {
						fields := strings.Fields(outsideString)
						insideContent.WriteString(fields[0])
						content = append(content, insideContent.String())
						if len(fields) >= 1 {
							content = append(content, fields[1:len(fields)]...)
						}
						insideContent.Reset()
						outsideContent.Reset()
					} else {
						insideContent.WriteString(outsideString)
						outsideContent.Reset()
					}
				}
			}
		} else {
			if inQuote {
				insideContent.WriteRune(r)
			} else {
				outsideContent.WriteRune(r)
			}
		}
	}

	if counter%2 != 0 {
		return []string{}, errors.New("Bad argument: unbalanced quotes")
	}

	outsideString := outsideContent.String()
	insideString := insideContent.String()
	//fmt.Printf("inside: %s\n", insideString)

	if !inQuote {
		if strings.HasPrefix(outsideString, " ") {
			if insideString != "" {
				content = append(content, insideString)
			}
			fields := strings.Fields(outsideString)
			content = append(content, fields...)
		} else if strings.Trim(outsideString, " ") != "" {
			fields := strings.Fields(outsideString)
			insideContent.WriteString(fields[0])
			content = append(content, insideContent.String())
			if len(fields) >= 1 {
				content = append(content, fields[1:len(fields)]...)
			}
		} else {
			if insideString != "" {
				content = append(content, insideString)
			}
		}
	}
	//fmt.Printf("content: %s\n", strings.Join(content, " "))

	return content, nil
}
