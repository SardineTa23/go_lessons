package main

import "fmt"

const Pi = 3.14

// 定数は型指定しない
const (
	Username = "kensho"
	Password = "pass"
)

func main() {
	fmt.Println(Pi, Username, Password)
}
