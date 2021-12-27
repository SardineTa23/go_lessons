package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)

	//DBなど
	//2021-12-27T23:19:00+09:00
	fmt.Println(t.Format(time.RFC3339))

	// 2021 December 27 23 20 15
	fmt.Println(t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}
