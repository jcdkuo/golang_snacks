package main

import "fmt"

type S struct {
	Name string
	Age  int
	//Address *int
}

func main() {
	a := S{
		Name: "aa",
		Age:  1,
		//Address: new(int),
	}
	b := S{
		Name: "aa",
		Age:  1,
		//Address: new(int),
	}

	fmt.Println(a == b)
}
