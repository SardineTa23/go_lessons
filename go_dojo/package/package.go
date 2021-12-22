package main

import (
	"fmt"
)

func main() {
	false := true
	if false {
		fmt.Println(false)
	}

	type int string
	var n int
	fmt.Println(n + "hoge")
}
