package main

import "fmt"

// Deprecated: test
func hello() {
	fmt.Println("hello world")
}

func returnError() error {
	return nil
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
}
