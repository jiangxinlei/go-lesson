package main

import (
	"fmt"
)

func enums() {
	const (
		cpp  = 0
		java = 1
		php  = 2
		js   = 3
	)

	// 可通过 iota 完成定义，自增值
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, java, php, js)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	enums()
}
