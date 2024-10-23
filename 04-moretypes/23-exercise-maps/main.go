package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	res := make(map[string]int)
	fields := strings.Fields(s)
	for _, field := range fields {
		res[field] += 1
	}

	return res
}

func main() {
	wc.Test(WordCount)
}
