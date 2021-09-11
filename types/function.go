package main

import "fmt"

func add(x int, y int) (int, int) {
	return x + y, x - y
}

func cal(price, item int) (result int) {
	result = price * item
	// 返り値の変数名を定義済みなので、return result 入らない
	return
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func foo(params ...int) {
	fmt.Println(len(params), params)

	for _, param := range params {
		fmt.Println(param)
	}
}

func main() {
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 2)
	fmt.Println(r3)

	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)

	func(x int) {
		fmt.Println("inner func", x)
	}(1)

	// クロージャー
	c1 := circleArea(3.14)
	fmt.Println(c1(3))
	fmt.Println(c1(2))

	// 可変長引数
	foo(100, 200)
	foo(100, 200, 300)

	s := []int{1, 2, 3}
	foo(s...)
}
