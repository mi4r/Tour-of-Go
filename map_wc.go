package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	lst := strings.Fields(s)
	res := make(map[string]int)
	for i := 0; i < len(lst); i++ {
		res[lst[i]]++
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
