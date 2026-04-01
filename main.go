package main

import (
	"fmt"
	"strconv"

	"github.com/mthaler/aparser"
)

func main() {
	fmt.Println("Hello, World!")

	var zeros []float64

	for {
		s := ""
		fmt.Scanln(s)
		if s == "" {
			break
		}

		z, err := strconv.ParseFloat(s, 64)

		zeros = append(zeros, z)
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("%+v\n", zeros)

	r, err := aparser.Eval("3 + 4")
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}
