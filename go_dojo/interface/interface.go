package main

import (
	"fmt"
)

// type Func func() string

// func (f Func) String() string { return f() }

// func main() {
// 	//無名関数をFunc型へキャストしている, 基底型から具象型へのキャストは可能
// 	var s fmt.Stringer = Func(func() string {
// 		return "hi"
// 	})
// 	fmt.Println(s)
// }

// d

// func main() {
// 	var i interface{}
// 	i = "fefef"
// 	switch v := i.(type) {
// 	case int:
// 		fmt.Println(v * 2)
// 	case string:
// 		fmt.Println(v + "hoge")
// 	default:
// 		fmt.Println("default")

// 	}
// }

type Hoge interface {
	M()
	N()
}
type fuga struct{ Hoge }

func (f fuga) M() {
	fmt.Println("Hi")
	f.Hoge.M()
}

func HiHoge(h Hoge) Hoge {
	return fuga{h}
}
