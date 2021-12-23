package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 100
	a[1] = 499
	fmt.Println(a)

	n := []int{1, 3, 5, 6, 9}
	fmt.Println(n)
	fmt.Println(n[2:5])

	var board = [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}

	fmt.Println(board)

	c := make([]int, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)
}
