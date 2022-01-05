package main

import (
	"fmt"
	"strings"
)

func main() {
	// input := ""
	input := "Apple Banana Apple Banana apple"
	fmt.Println(wordCount(input))

}

func wordCount(s string) map[string]int {
	m := map[string]int{}
	if s == "" {
		return m
	}
	new := strings.Split(s, " ")
	for i := 0; i < len(new); i++ {
		m[new[i]] = m[new[i]] + 1
	}
	return m
}
