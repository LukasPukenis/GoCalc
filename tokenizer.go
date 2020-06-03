package main

import (
	"regexp"
)

func tokenize(input string) ([]string, error) {
	re, err := regexp.Compile(`(\d+|[*+-/\(\)]|sin|cos)`)
	if err != nil {
		return nil, err
	}

	matches := re.FindAll([]byte(input), -1)

	result := []string{}
	for _, match := range matches {
		result = append(result, string(match))
	}

	return result, err
}
