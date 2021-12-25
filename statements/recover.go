package main

import "fmt"

func connectDB() {
	panic("unable to connect database!")
}

func save() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()

	// panicより先にrecoverを書いておく必要がある。
	connectDB()
}

func main() {
	save()
	fmt.Println("OK?")
}
