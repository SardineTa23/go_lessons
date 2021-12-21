// メソッドはレシーバと紐付けられた関数
package main

type MyInt int

// 動かない
// func (n MyInt) Inc() { n++ }

// func main() {
// 	var n MyInt
// 	println(n)
// 	n.Inc()
// 	println(n)
// }

// *MyInt  ポインタの指す中身をレシーバに渡す
func (n *MyInt) Inc() { *n++ }

func main() {
	var n MyInt
	println(n)
	n.Inc()
	println(n)
}
