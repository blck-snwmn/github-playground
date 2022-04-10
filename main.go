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
	returnError()
}
