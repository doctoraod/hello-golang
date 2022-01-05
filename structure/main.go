package main

import "fmt"

type Book struct {
	Name   string
	Author string
}

func main() {
	book := Book{
		Name:   "Harry Potter",
		Author: "J. K. Rowling",
	}
	fmt.Println(book)
}

func (b Book) String() string {
	return b.Name + " by " + b.Author
}
