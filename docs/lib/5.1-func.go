package main

import (
	"fmt"
)

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operation: " + op)
	}
}

// 返回多个值 13 / 3 = ... 1，
func div(a, b int) (int, int) {
	return a / b, a % b
}

// q即第一个返回值，r是第2个
func div1(a, b int) (q, r int) {
	return a / b, a % b
}

func main() {
	fmt.Println(eval(3, 4, "*"))

	fmt.Println(div(13, 3))

	q, r := div1(13, 3)
	q, _ := div1(13, 3) // 只返回第一个返回值
	fmt.Println(q, r)
}
