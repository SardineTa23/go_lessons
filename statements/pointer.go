package main

import "fmt"

func one(x *int) {
	*x = 1
}

func main() {
	var n int = 100
	// one関数の引数はポインタ型なので、ポインタアドレスを渡さないといけない
	one(&n)
	fmt.Println(n)

	/*
		var n int = 100
		fmt.Println(n)

		// メモリアドレス取得
		fmt.Println(&n)

		// ポインタ型（ポインタ変数
		var p *int = &n
		fmt.Println(p)

		// メモリアドレスから値取得
		fmt.Println(*p)
	*/
}
