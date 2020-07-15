package main

import "fmt"

var aa = 3
var bb = 4
var cc = "hello" // 这些变量都是main包内部的变量，go没有全局变量说法， := 声明只能在函数内使用

// 上面也可以这么写
var (
	dd = 5
	ee = 6
)

// 声明变量和类型
func varZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

// 声明变量和类型并赋值
func varInitValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

// 声明变量并赋值，打印时自动推断变量类型
func varTypeValue() {
	var a, b, c, s = 3, 4, true, "abc"
	fmt.Println(a, b, c, s)
}

// 声明变量并赋值，打印时自动推断变量类型，:= 同 var
func varShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5 // 可再次声明，并覆盖第一次声明的值
	fmt.Println(a, b, c, s)
}

func main() {
	varZeroValue()
	varInitValue()
	varTypeValue()
	varShorter()
}
