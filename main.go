package main

import (
	"fmt"
)

//Deprecated: test
type Hoge struct{}

//Deprecated: test
func hello() {
	fmt.Println("hello world")
}

// Deprecated: aaaaa
const x = 10

//Deprecated: aaaa
func getId() {
	fmt.Println(x)
}

// func returnError() error {
// 	return errors.New(fmt.Sprintf("aaa"))
// }

func main() {
	hello()
	hello()
	x := 10
	if x%1 == 0 {
		fmt.Println("hello")
	}
	xs := fmt.Sprintf("%v", x)
	fmt.Println(xs)
	getId()
	h := Hoge{}
	fmt.Println(h)
}
