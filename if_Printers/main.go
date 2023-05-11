package main

import "fmt"

// 定義一個通用的打印器（Printer）接口
type Printer interface {
	Print()
}

// 定義一個結構體 Person，實現了 Printer 接口
type Person struct {
	Name string
	Age  int
}

// 實現 Printer 接口中的 Print 方法
func (p Person) Print() {
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}

// 定義一個結構體 Book，實現了 Printer 接口
type Book struct {
	Title  string
	Author string
}

// 實現 Printer 接口中的 Print 方法
func (b Book) Print() {
	fmt.Printf("Title: %s, Author: %s\n", b.Title, b.Author)
}

func main() {
	// 創建一個 Printer 的切片，其中包含了一個 Person 和一本 Book
	printers := []Printer{
		Person{Name: "Alice", Age: 30},
		Book{Title: "Golang in Action", Author: "William Kennedy"},
	}

	// 遍歷切片，呼叫 Print 方法
	for _, printer := range printers {
		printer.Print()
	}
}
