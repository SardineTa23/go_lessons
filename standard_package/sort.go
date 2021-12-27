package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{5, 3, 2, 8, 7}
	s := []string{"d", "a", "f"}
	p := []struct {
		Name string
		Age  int
	}{
		{"Nancy", 20},
		{"Vera", 40},
		{"Mike", 30},
		{"Bob", 50},
	}
	// [5 3 2 8 7] [d a f] [{Nancy 20} {Vera 40} {Mike 30} {Bob 50}]
	fmt.Println(i, s, p)

	sort.Ints(i)
	sort.Strings(s)
	sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })

	// [2 3 5 7 8] [a d f] [{Bob 50} {Mike 30} {Nancy 20} {Vera 40}]
	fmt.Println(i, s, p)
}
