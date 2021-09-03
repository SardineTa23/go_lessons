package main

import "fmt"

func main() {
	var (
		u8 uint8 = 255
		i8 int8  = 127
	)

	fmt.Println(u8, i8)
	fmt.Printf("type=%T value=%v", u8, u8)

	// x := 1 + 1
	// fmt.Println(x)
	// fmt.Println("2 + 3 =", 2+3)

	// x := 0
	// fmt.Println(x)
	// x++
	// fmt.Println(x)

}
