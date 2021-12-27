package main

import "fmt"

const (
	c1 = iota
	c2
	c3
)

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

func main() {
	//0 1 2
	fmt.Println(c1, c2, c3)

	//1024 1048576 1073741824
	fmt.Println(KB, MB, GB)
}
