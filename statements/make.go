package main

import "fmt"

func main() {
	/*
		// newすると、メモリは確保される
		var n *int = new(int)

		// 初期値の0が入っている
		fmt.Println(*n)
		fmt.Println(&n)

		var p2 *int
		fmt.Println(p2)
	*/
	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	ch := make(chan int)
	fmt.Printf("%T\n", ch)

	var p *int = new(int)
	fmt.Printf("%T\n", p)

	var st = new(struct{})
	fmt.Printf("%T\n", st)
}
