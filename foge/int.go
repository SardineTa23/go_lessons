package main

import "fmt"

func main() {
	var (
		u8 uint8 = 255
	)

	fmt.Println(u8)
	fmt.Printf("%T %v", u8, u8)

	x := 0
	fmt.Println(x)
	x++
	fmt.Println(x)
}
