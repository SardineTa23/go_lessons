package main

import (
	"encoding/json"
	"fmt"
)

type T struct{}

// encode(Marshal)するときのキーを型の後ろに付ける
type Person struct {
	Name      string   `json:"name,omitempty"`
	Age       int      `Json:"age,omitempty"`
	Nicknames []string `json:"nicknames"`
	T         *T       `json:"T,omitempty"`
}

func (p *Person) UnmarshalJSON(b []byte) error {
	type Person2 struct {
		Name string
	}
	var p2 Person2
	err := json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Println(err)
	}
	p.Name = p2.Name + "!"
	return err
}

func (p Person) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr." + p.Name,
	})
	return v, err
}

func main() {
	b := []byte(`{"name":"Mike", "age": 20, "nicknames": ["a", "b", "c"]}`)
	var p Person

	// jsonをtypeにいれてくれる
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	//json_encode的なやつ
	v, _ := json.Marshal(p)
	fmt.Println(string(v))
}
