package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hello")
	fmt.Println("hello" + "world")
	fmt.Println(string("helloworld"[0]))

	var s string = "Hello world"
	fmt.Println(strings.Replace(s, "H", "x", 1))

	fmt.Println(`Test
Test`)
}
