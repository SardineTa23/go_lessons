package main

import "fmt"

//何でも受け付ける
func do(i interface{}) {
	/*
		// タイプアサーション, intに直す
		ii := i.(int)
		fmt.Println(ii)
	*/

	/*
		ss := i.(string)
		fmt.Println(ss + "!")
	*/

	// switch-type文, 型によって処理を分ける
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know %T\n", v)
	}
}

func main() {
	// // iは代入時点ではintではなくinterface
	// var i interface{} = 10
	// do(i)

	do(10)
	do("Mike")
	do(true)

	var i int = 10
	ii := float64(10)
	fmt.Println(i, ii)
}
