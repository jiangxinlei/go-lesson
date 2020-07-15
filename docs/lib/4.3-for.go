package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 整数转成二进制
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 死循环
func con() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(convertToBin(13))

	printFile("abc.txt")
}
