package main

import (
	"fmt"

	"github.com/mthaler/aparser"
)

func main() {
	fmt.Println("Hello, World!")

	r, err := aparser.Eval("3 + 4")
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}
