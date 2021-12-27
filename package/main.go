package main

import (
	"fmt"
	"package/mylib"
	"package/mylib/under"

	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))

	mylib.Say()
	under.Hello()
	person := mylib.Person{Name: "Mike", Age: 20}
	fmt.Println(person)

	fmt.Println(mylib.Public)

	spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", quote.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
}
