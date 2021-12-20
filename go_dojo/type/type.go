package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	var sum int
// 	sum = 5 + 6 + 3

// 	// float64でキャスト
// 	avg := float64(sum) / 3
// 	if avg > 4.5 {
// 		println("good")
// 	}
// }

// // 要素がintのスライス
// var ns []int

// // 配列は型と要素数を指定,要素数も型の一部
// var arr [5]int

// // キーがstring,valueがintのmap
// var m map[string]int

// var p struct {
// 	name string
// 	age  int
// }

// func main() {
// 	a := []int{10, 20}
// 	fmt.Println(a, cap(a))

// 	b := append(a, 30)

// 	// コピーが作成される挙動をするappendを使ってbを定義したあとにaを書き換えても、bの値には影響を及ぼさない
// 	a[0] = 100
// 	fmt.Println(b, cap(b))

// 	c := append(b, 40)
// 	b[1] = 200
// 	fmt.Println(c, cap(c))
// }

// func main() {
// 	var ns [10000000]int
// 	ms := ns[:10]
// 	fmt.Println(len(ms), cap(ms))
// 	// 結果 10 10000000
// 	// capがでかすぎるから、スライス作成の配列コピー時にトリミングしてあげよう

// 	mss := ns[:10:10]
// 	fmt.Println(len(mss), cap(mss))
// 	// 結果 10 10
// 	// capが縮んだ
// }

func main() {
	counts := map[string]int{}

	str := "dog dog dog cat dog"
	for _, s := range strings.Split(str, " ") {
		counts[s]++
	}
	fmt.Println(counts)
	// 出力 map[cat:1 dog:4]
}
