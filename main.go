package main

import (
	"errors"
	"fmt"
)

//Deprecated: test
func hello() {
	fmt.Println("hello world")
}

//Deprecated: aaaa
func getId() {

}

func returnError() error {
	return errors.New(fmt.Sprintf("aaa"))
}

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
}
