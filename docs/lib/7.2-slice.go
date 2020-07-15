package main

import (
	"fmt"
)

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6] =", arr[2:6]) // [2, 3, 4, 5]
	fmt.Println("arr[:6] =", arr[:6])   // [0, 1, 2, 3, 4, 5]
	fmt.Println("arr[2:] =", arr[2:])   // [2, 3, 4, 5, 6, 7]
	fmt.Println("arr[:] =", arr[:])     // [0, 1, 2, 3, 4, 5, 6, 7]

	// 切片不是值类型，可以实现引用类型的效果
	s1 := arr[2:]
	fmt.Println("s1 =", s1)

	s2 := arr[:]
	fmt.Println("s2 =", s2)

	updateSlice(s1)
	fmt.Println("after alice s1 =", s1)
	fmt.Println(arr)

	updateSlice(s2)
	fmt.Println("after alice s2 =", s2)
	fmt.Println(arr)

	fmt.Println("--------------")

	fmt.Println("s2 =", s2)
	s2 = s2[:5]
	fmt.Println("s2 =", s2)

	s2 = s2[2:]
	fmt.Println("s2 =", s2)
	fmt.Println(arr)

	fmt.Println("--------------")

	arr1 := [...]int{3, 4, 1, 8, 0, 9, 7}
	s3 := arr1[2:6]
	s4 := s3[3:5]
	fmt.Println("s3 =", s3)
	fmt.Println("s4 =", s4)
}
