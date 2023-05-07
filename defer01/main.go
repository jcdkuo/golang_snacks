package main

import "fmt"

func c() int {
	i := 3
	defer func() { i++ }()
	return 1
}

func main() {
	language := "Go"
	defer fmt.Print(language + "\n")

	language = "Node.js"
	fmt.Print("Hello ")

	language = "Go"
	defer fmt.Print(" to " + language + "\n")

	language = "Node.js"
	defer fmt.Print("from " + language)
	fmt.Print("Hello ")

	ret := c()
	fmt.Println("c = ", ret)
}
