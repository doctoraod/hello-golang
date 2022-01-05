package main

import "fmt"

func varia(s ...string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

func main() {
	s := []string{"a", "b"}
	varia(s...)
}
