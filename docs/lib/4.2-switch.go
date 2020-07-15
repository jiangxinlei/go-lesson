package main

import "fmt"

// 参数是 score，为int类型，返回值是 string 类型
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}

	return g
}

func main() {
	fmt.Println(grade(70), grade(90))
}
