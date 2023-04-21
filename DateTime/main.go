package main

import (
	"fmt"
	"time"
)

func main() {
	// 獲取當前時間
	now := time.Now()

	// 將時間格式化為 2006-01-02 15:04:05 格式的字符串
	formattedTime := now.Format("2006-01-02 15:04:05")

	fmt.Println(formattedTime) // 輸出: 2023-04-20 15:30:45
}
