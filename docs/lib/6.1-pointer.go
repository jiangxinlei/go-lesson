package main

import (
	"fmt"
)

// 利用指针交换a, b值
func swap(a, b *int) {
	*a, *b = *b, *a
}

// 也可以这样交换a, b值
func swap1(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b) // 4, 3

	a, b = swap1(a, b)
	fmt.Println(a, b) // 4, 3
}
