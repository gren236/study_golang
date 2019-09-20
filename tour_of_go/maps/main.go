package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	stringArray := strings.Fields(s)
	result := map[string]int{}
	for _, v := range stringArray {
		_, ok := result[v]
		if ok {
			result[v] = result[v] + 1
		} else {
			result[v] = 1
		}
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
