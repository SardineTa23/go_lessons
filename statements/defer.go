package main

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("world foo")
	fmt.Println("hello foo")
}

// deferは関数の処理が終わったあとに実行される
func main() {
	// foo()
	// defer fmt.Println("world")
	// fmt.Println("hello")

	// fmt.Println("run")
	// // 3, 2, 1の順で実行される
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// fmt.Println("success")

	file, _ := os.Open("./defer.go")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}
