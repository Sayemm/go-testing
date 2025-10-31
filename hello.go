package main

import "fmt"

func Hello() string {
	return "Hello World"
}

func main() {
	fmt.Println("Hello World")
}

/*
It is good to separate the "domain" code from the outside world (side-effects).
The fmt.Println is a side effect (printing to stdout), and the string we send in is our domain.
*/
