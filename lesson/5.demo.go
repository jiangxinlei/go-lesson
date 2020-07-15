package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "everyone", "ths greeting object")
	flag.Parse()
	fmt.Printf("hello, %v!\n", name)
}
