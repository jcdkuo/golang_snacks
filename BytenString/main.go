package main

import "fmt"

func main() {
	str := "Hello, World!"

	// 將字符串轉換為 []byte
	byteSlice := []byte(str)

	// 輸出 []byte 的內容
	fmt.Printf("%v\n", byteSlice) // 輸出: [72 101 108 108 111 44 32 87 111 114 108 100 33]

	user := "jerry"

	b2 := []byte("user:" + user)

	fmt.Printf("%v\n", b2)

}
