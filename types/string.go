package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println("hello, world")
	// fmt.Println(string("hello, world"[0]))

	// 文字置換
	var s string = "Hello World"
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)
	fmt.Println(strings.Contains(s, "World"))
	fmt.Println(strings.Contains(s, "hoge"))

	fmt.Println("Test\n" + "Test")
	fmt.Println(`Test
	Test
	                     Test`)

	fmt.Println(`"`)
}
