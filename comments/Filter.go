package main

import (
	"bufio"
	"os"
	"strings"
)

var filtered []string

func init() {
	file, _ := os.Open("word-filter.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		filtered = append(filtered, scanner.Text())
	}
}

func filter(text string) string {
	for _, v := range filtered {
		for {
			i := strings.LastIndex(strings.ToLower(text), strings.ToLower(v))

			if i == -1 {
				break
			} else {
				text = string(text[0:i]) + strings.Repeat("*", len(v)) + string(text[i+len(v):])
			}
		}
	}

	return text
}
