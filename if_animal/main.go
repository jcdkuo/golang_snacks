package main

import "fmt"

// 定義一個名為 Animal 的 interface
type Animal interface {
	Speak() string
}

// 定義一個名為 Dog 的結構體，實現了 Animal interface
type Dog struct{}

// 實現 Animal interface 中的 Speak 方法
func (d Dog) Speak() string {
	return "Woof!"
}

// 定義一個名為 Cat 的結構體，實現了 Animal interface
type Cat struct{}

// 實現 Animal interface 中的 Speak 方法
func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	// 創建一個 Animal 的切片，其中包含了一隻狗和一隻貓
	//animals := []Animal{Dog{}, Cat{}}
	animals := []Animal{Cat{}, Dog{}, Cat{}}

	// 遍歷切片，並呼叫 Speak 方法
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
