package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// 出力が変わる
// {Mike 22} -> My name is Mike
func (p Person) String() string {
	return fmt.Sprintf("My name is %v", p.Name)
}

type UserNotFound struct {
	Username string
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found: %v", e.Username)
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{Username: "mike"}
}

func main() {
	mike := Person{"Mike", 22}
	fmt.Println(mike)

	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}
