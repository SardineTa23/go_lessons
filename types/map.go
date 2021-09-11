package main

import "fmt"

func main() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Print(m["apple"])
	m["new"] = 300
	fmt.Println(m)

	v, ok := m["apple"]
	fmt.Println(v, ok)

	m2 := make(map[string]int)
	fmt.Println(m2)
	m2["pc"] = 5000
	fmt.Println(m2)
}
