package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"
)

func main() {
	var n int
	if len(os.Args) > 1 {
		arg := os.Args[1]
		n, _ = strconv.Atoi(arg)
	}
	if n == 0 {
		n = 1
	}
	for i := 0; i < n; i++ {
		fmt.Println(uuid.NewString())
	}
}
