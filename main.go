package main

import (
	"fmt"
)

func main() {
	// Exercise function
	var message1 string = greeting("Aod")
	fmt.Println(message1)
	var message2 string = greetingWithAge("Aod", 30)
	fmt.Println(message2)
	var message3 string = greetingWithAgeAndDrink("Aod", 30, "Pepsi")
	fmt.Println(message3)
	// Exercise conditions
	var result1 bool = isOdd(10)
	fmt.Println(result1)

	// Exercise loops
	var resultLoop1 int = sumOfFirst(3)
	fmt.Println(resultLoop1) // result 6
	var resultLoop2 bool = isPrime(1)
	fmt.Println(resultLoop2) // result false

	// Pointer
	var d int = 2
	double(&d)
	fmt.Println(d) // result 4
	name := "Aod"
	appendGreeting(&name)
	fmt.Println(name)
}

func greeting(name string) string {
	return "Hello, " + name
}

// greetingWithAge("Pallat", 30) should return "Hello, Pallat. You are 30 years old."
func greetingWithAge(name string, age uint) string {
	var result string = fmt.Sprintf("Hello, %s. You are %d years old.", name, age)
	return result
}

// // greetingWithAgeAndDrink("Pallat", 30, "Cola") should return "Hello, Pallat. You are 30 years old and your favorite drink is Cola."
func greetingWithAgeAndDrink(name string, age uint, drink string) string {
	var result string = fmt.Sprintf("Hello, %s. You are %d years old and your favorite drink is %s.", name, age, drink)
	return result
}

func isOdd(n int) bool {
	return n%2 != 0
}

func sumOfFirst(n int) int {
	var result int = 0
	for i := 1; i <= n; i++ {
		result = result + i
	}
	return result
}

// A prime number is a number greater than 1 with only two factors – themselves and 1
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func double(d *int) {
	*d = *d * 2
}

func appendGreeting(s *string) {
	*s = fmt.Sprintf("Hi, %s", *s)
}
