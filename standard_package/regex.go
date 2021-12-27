package main

import (
	"fmt"
	"regexp"
)

func main() {
	//正規表現マッチ
	match, _ := regexp.MatchString("a([a-z+]+)e", "apple")
	fmt.Println(match)

	r := regexp.MustCompile("a([a-z+]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	// s:= "/view/test"
	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-z0-9]+)$")
	fs := r2.FindString("/view/test")
	fmt.Println(fs)

	// [/view/test view test]
	fss := r2.FindStringSubmatch("/view/test")
	fmt.Println(fss)
	// [/save/test save test]
	fss = r2.FindStringSubmatch("/save/test")
	fmt.Println(fss)
}
