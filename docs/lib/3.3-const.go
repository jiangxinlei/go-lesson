package main

import (
	"fmt"
	"math"
)

const a = 3

func consts() {
	const b = 4
	var c int

	c = int(math.Sqrt(a*a + b*b))

	fmt.Println(c)
}

func main() {
	consts()
}
