package main

import "fmt"

func main() {
	f := 1.11
	ff := int(f)
	fmt.Printf("%T %v %d", ff, ff, ff)

	m := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
	fmt.Printf("%T %v", m, m)
}
