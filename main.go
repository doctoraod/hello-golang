package main

import (
	"fmt"
)

func main() {
	var message1 string = greeting("Aod")
	fmt.Println(message1)
	var message2 string = greetingWithAge("Aod", 30)
	fmt.Println(message2)
	var message3 string = greetingWithAgeAndDrink("Aod", 30, "Pepsi")
	fmt.Println(message3)
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
