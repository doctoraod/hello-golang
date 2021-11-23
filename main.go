package main

import "fmt"

func main() {
	var message string = greeting("Aod", "Suvichan")
	fmt.Println(message)
	var num int = add(1, 2)
	fmt.Println(num)
	const (
		sunday = iota
		monday
	)
	fmt.Println(sunday)
	fmt.Println(monday)
}

func greeting(firstName, lastName string) string {
	return "Hello, " + firstName + lastName
}

func add(a, b int) int {
	return a + b
}
