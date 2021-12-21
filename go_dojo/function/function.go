package main

import "fmt"

// func main() {
// 	// msg := "Hello, world"

// 	// func() {
// 	// 	println(msg)
// 	// }()
// 	// //()ですぐ呼び出し

// 	// fs := make([]func(), 3)
// 	// for i := range fs {
// 	// 	i := i
// 	// 	fs[i] = func() { fmt.Println(i) }
// 	// }

// 	// for _, f := range fs {
// 	// 	f()
// 	// }

// 	// n := 10
// 	// m := n
// 	// n = 20
// 	// fmt.Println(n, m)

// 	// p := struct {
// 	// 	age  int
// 	// 	name string
// 	// }{age: 10, name: "Gopher"}

// 	// p2 := p
// 	// p2.age = 20

// 	// println(p.age, p.name)
// 	// println(p2.age, p2.name)
// }

func f(xp *int) {
	// 渡されてきたポインタの位置に100を代入
	// *xp = 100

	var y int
	xp = &y
}

func main() {
	var x int

	// &xは、xのポインタの値を取得
	var xp *int = &x

	f(xp)
	*xp = 100
	fmt.Println(x)
}
