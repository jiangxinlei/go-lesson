package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func euler() {
	// c := 3 + 4i
	// fmt.Println(cmplx.Abs(c))

	// fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)

	fmt.Println("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}
func triangle() {
	a, b := 3, 4
	var c int

	c = int(math.Sqrt(float64(a*a + b*b)))

	fmt.Println(c)
}
func main() {
	fmt.Println("hello world")
	euler()
	triangle()
}
