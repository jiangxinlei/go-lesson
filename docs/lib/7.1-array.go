package main

import (
	"fmt"
)

func printArray(arr *[5]int) {
	arr[0] = 100

	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}

	// 让编译器去数有多少
	arr3 := [...]int{2, 4, 6, 8, 10}

	// 定义多维数组，4行5列
	var arr4 [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(arr4)

	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	// 用range关键字获取元素下标和值，如果只要值，就将下标 i 写成 _
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	// 函数中的参数是 [5]int，所以传参时也必须是 [5]int
	// 使用指针时这里就需要取地址
	printArray(&arr1)
	printArray(&arr3)

	// 数组是值类型，所以即使在函数中改变了 arr[0] 的值，实际上数组还是没有改变
	// 可以用指针改变
	fmt.Println(arr1, arr3)
}
