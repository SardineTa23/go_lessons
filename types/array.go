package main

import "fmt"

func main() {
	// 配列はサイズと型両方指定する。
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	// var b [2]int = [2]int{100, 200}
	// fmt.Println(b)
	// fmt.Printf("%T", b)

	// 数字を入れないとスライス型になり、要素数を変更できる
	var b []int = []int{100, 200}
	b = append(b, 300)
	fmt.Println(b)
}
