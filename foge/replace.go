package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 1
	xx := float64(x)
	fmt.Println(xx)

	var y float64 = 1.2
	yy := int(y)
	fmt.Println(yy)

	var s string = "134"

	// _ は使わない
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v", i, i)

	h := "Hello World"
	fmt.Println(string(h[0]))
}
