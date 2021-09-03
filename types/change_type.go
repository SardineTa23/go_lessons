package main

import (
	"fmt"
	"strconv"
)

// 型変換
func main() {

	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	var s string = "14"
	// Atoiの返り値が２つある、そのため _ を書いて不要な返り値を受け取る
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v\n", i, i)

	h := "Hello World"
	fmt.Printf("%T, %v\n", h, h)
	fmt.Println(string(h[0]))
}
