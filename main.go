package main

import (
	"fmt"
)

// Add 函数用于计算两个整数的和
func Add(a, b int) int {
	return a + b
}

func main() {
	result := Add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)
}
