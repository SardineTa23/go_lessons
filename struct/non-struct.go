package main

import "fmt"

type MyInt int

func (i MyInt) Double() int {
	return int(i * 2)
}

func main() {
	myiInt := MyInt(10)
	fmt.Println(myiInt.Double())
}
