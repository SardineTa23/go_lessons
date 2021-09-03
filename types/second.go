package main

import "fmt"

func main() {
	// var i int = 1
	// var f64 float64 = 1.2
	// var s string = "test"
	// var t, f bool = true, false

	// var (
	// 	i    int
	// 	f64  float64
	// 	s    string
	// 	t, f bool
	// )
	// fmt.Println(i, f64, s, t, f)

	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	// 型表示
	fmt.Printf("%T\n", xi)
	fmt.Printf("%T\n", xs)
}
