package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "abc.txt"

	/* constents, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", constents)
	} */

	// 上述也可这么写
	if constents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", constents)
	}
}
