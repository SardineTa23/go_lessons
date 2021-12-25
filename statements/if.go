package main

import "fmt"

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func main() {
	result := by2(10)
	if result == "ok" {
		fmt.Println("OKOK")
	}

	// 一行で簡略の書き方
	if result2 := by2(6); result2 == "ok" {
		fmt.Println("result2 ok")
	}

	num := 4

	if num%2 == 0 {
		fmt.Println("if")
	} else if num%3 == 0 {
		fmt.Println("else if")
	} else {
		fmt.Println("else")
	}
}
